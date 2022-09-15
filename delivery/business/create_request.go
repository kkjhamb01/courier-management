package business

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/kkjhamb01/courier-management/delivery/db"
	"github.com/kkjhamb01/courier-management/delivery/model"
	"github.com/kkjhamb01/courier-management/delivery/services"
	storage "github.com/kkjhamb01/courier-management/delivery/storage/lock"
	commonPb "github.com/kkjhamb01/courier-management/grpc/common/go"
	deliveryPb "github.com/kkjhamb01/courier-management/grpc/delivery/go"
	pricingPb "github.com/kkjhamb01/courier-management/grpc/pricing/go"
	"github.com/kkjhamb01/courier-management/party/proto"
	uaa "github.com/kkjhamb01/courier-management/uaa/proto"
	"google.golang.org/grpc"
)

func CreateRequest(ctx context.Context, customerId string, vehicleType commonPb.VehicleType,
	origin *deliveryPb.CreateRequestLocation, destinations []*deliveryPb.CreateRequestLocation, schedule *time.Time, requiredWorkers int32) (deliveryPb.Request, error) {

	logger.Infof("CreateRequest customerId = %v, vehicleType = %v, origin = %+v, destinations = %+v, schedule = %v, requiredWorkers = %v", customerId, vehicleType, origin, destinations, schedule, requiredWorkers)

	// check if the user is active (not banned, etc)
	conn, err := services.ConnectToParty()
	if err != nil {
		logger.Error("failed to connect to the party service", err)
		return deliveryPb.Request{}, uaa.Internal.Error(err)
	}
	GetClientUserStatusResponse, err := proto.NewUserStatusServiceClient(conn).GetClientUserStatus(ctx, &proto.GetClientUserStatusRequest{
		UserId: customerId,
	})
	if err != nil {
		logger.Error("failed to get user account to check the status", err)
		return deliveryPb.Request{}, uaa.Internal.Error(err)
	}

	logger.Infof("CreateRequest GetClientUserStatusResponse = %%v", GetClientUserStatusResponse)

	// TODO: @Majid: please remind me to talk about these status values
	if GetClientUserStatusResponse.Status != proto.UserStatus_USER_STATUS_AVAILABLE {
		err = fmt.Errorf("the user who creating the request is not available. user status: %v", GetClientUserStatusResponse.Status.String())
		logger.Error("the user creating the request is not available", err)
		return deliveryPb.Request{}, uaa.InvalidArgument.Error(err)
	}

	tx := db.MariaDbClient().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	// check if schedule has been set correctly
	if schedule != nil && schedule.Before(time.Now()) {
		err = fmt.Errorf("the supplied schedule: %v is for the past, now: %v", schedule, time.Now())
		logger.Error("the supplied schedule is not valid", err)
		return deliveryPb.Request{}, uaa.InvalidArgument.Error(err)
	}
	requestStatus := deliveryPb.RequestStatus_PROCESSING
	if schedule != nil {
		requestStatus = deliveryPb.RequestStatus_SCHEDULED
	}

	// check if requiredWorkers is set only when the vehicle type is VAN
	if requiredWorkers > 0 && vehicleType != commonPb.VehicleType_SMALL_VAN &&
		vehicleType != commonPb.VehicleType_MEDIUM_VAN && vehicleType != commonPb.VehicleType_LARGE_VAN {
		err = errors.New("helping workers are only available for VANs")
		logger.Error("the supplied requiredWorkers is not valid", err)
		return deliveryPb.Request{}, uaa.InvalidArgument.Error(err)
	}

	logger.Infof("CreateRequest check isOriginAndDestinationWithinAllowedCities")

	ok, err := isOriginAndDestinationWithinAllowedCities(ctx, origin, destinations)
	if err != nil {
		logger.Error("failed to check whether the cities are within valid geo locations (cities)", err)
		return deliveryPb.Request{}, uaa.Internal.Error(err)
	}

	logger.Infof("CreateRequest check isOriginAndDestinationWithinAllowedCities = %%v", ok)

	if !ok {
		err = errors.New("the request is failed due to the invalid origin/destination location")
		logger.Error("locations are not valid", err)
		return deliveryPb.Request{}, uaa.InvalidLocationOrigin.Error(err)
	}

	locations := make([]model.RequestLocation, 0, len(destinations)+1)

	destinationsModel := make([]model.RequestLocation, len(destinations))
	for i, destination := range destinations {
		destinationsModel[i] = model.RequestLocation{
			Lat:      destination.Lat,
			Lon:      destination.Lon,
			Order:    i + 1,
			IsOrigin: false,
		}

		if destination.FullName != nil {
			destinationsModel[i].FullName = destination.FullName.Value
		}
		if destination.PhoneNumber != nil {
			destinationsModel[i].PhoneNumber = destination.PhoneNumber.Value
		}
		if destination.AddressDetails != nil {
			destinationsModel[i].AddressDetails = &destination.AddressDetails.Value
		}
		if destination.CourierInstructions != nil {
			destinationsModel[i].CourierInstructions = &destination.CourierInstructions.Value
		}
	}
	locations = append(locations, destinationsModel...)

	originModel := model.RequestLocation{
		Lat:      origin.Lat,
		Lon:      origin.Lon,
		IsOrigin: true,
	}
	if origin.FullName != nil {
		originModel.FullName = origin.FullName.Value
	}
	if origin.PhoneNumber != nil {
		originModel.PhoneNumber = origin.PhoneNumber.Value
	}
	if origin.AddressDetails != nil {
		originModel.AddressDetails = &origin.AddressDetails.Value
	}
	if origin.CourierInstructions != nil {
		originModel.CourierInstructions = &origin.CourierInstructions.Value
	}
	locations = append(locations, originModel)

	logger.Infof("CreateRequest start calculate price origin = %%v, destinations = %%v", origin, destinations)

	price, err := CalculateCourierPrice(customerId, origin,
		destinations,
		vehicleType,
		requiredWorkers)

	if err != nil {
		logger.Error("failed to get price", err)
		return deliveryPb.Request{}, uaa.Internal.Error(err)
	}

	requestModel := model.Request{
		CustomerId:             customerId,
		Locations:              locations,
		EstimatedDuration:      price.EstimatedDuration,
		EstimatedDistanceMeter: int(price.EstimatedDistance),
		FinalPrice:             price.Amount,
		FinalPriceCurrency:     price.Currency,
		VehicleType:            vehicleType.String(),
		RequiredWorkers:        requiredWorkers,
		HumanReadableId:        requestHumanReadableId(),
		ScheduleOn:             schedule,
		Status:                 requestStatus.String(),
	}

	// store request
	result := tx.Create(&requestModel)
	if result.Error != nil {
		logger.Error("failed to create request", result.Error)
		return deliveryPb.Request{}, uaa.Internal.Error(result.Error)
	}

	logger.Infof("CreateRequest init lock")

	err = storage.InitLockRequest(ctx, requestModel.ID)
	if err != nil {
		logger.Error("failed to init request lock", err)
		return deliveryPb.Request{}, uaa.Internal.Error(err)
	}

	// publish new request event if the request does not have a schedule
	if schedule == nil {
		err = publishNewRequestEventOnMessaging(ctx, requestModel.ToProto())
		if err != nil {
			logger.Error("failed to publish new request event", err)
			return deliveryPb.Request{}, uaa.Internal.Error(err)
		}
	} else {
		logger.Infof("the request %v has been scheduled to run on %v, current time: %v", requestModel.ID, requestModel.ScheduleOn, time.Now())
	}

	var originId string
	for _, location := range requestModel.Locations {
		if location.IsOrigin {
			originId = location.ID
		}
	}

	rideStatus := model.RideStatus{
		RideLocationId:   originId,
		RequestId:        requestModel.ID,
		Status:           requestModel.Status,
		CancellationNote: "",
	}
	err = tx.Create(&rideStatus).Error
	if err != nil {
		logger.Error("failed to create ride status", err)
		return deliveryPb.Request{}, uaa.Internal.Error(err)
	}

	return requestModel.ToProto(), nil
}

func CreateRequestSchedulerJob() {
	logger.Info("Create Request Job Scheduler Started ...")
	var scheduledRequests []*model.Request

	result := db.MariaDbClient().
		Preload("Locations").
		Where("status = ? AND schedule_on < ?", deliveryPb.RequestStatus_SCHEDULED.String(), time.Now()).
		Find(&scheduledRequests)
	if result.Error != nil {
		logger.Error("failed to list addresses", result.Error)
		return
	}

	for _, request := range scheduledRequests {
		//logger.Infof("starting the scheduled request: %v", request)

		tx := db.MariaDbClient().Begin()

		lock, err := storage.LockRequest(context.Background(), request.ID)
		if err != nil {
			tx.Rollback()
			//logger.Error("failed to acquire request lock", err)
			continue
		}

		err = tx.Model(&request).
			Updates(map[string]interface{}{
				"status": deliveryPb.RequestStatus_PROCESSING.String(),
			}).Error
		if err != nil {
			tx.Rollback()
			logger.Error("failed to update the model (from scheduled) to processing", err)
			continue
		}

		err = publishNewRequestEventOnMessaging(context.Background(), request.ToProto())
		if err != nil {
			tx.Rollback()
			logger.Error("failed to publish new request created", err, tag.Str("requestId", request.ID))
			continue
		}

		var originLocation model.RequestLocation
		for _, location := range request.Locations {
			if location.IsOrigin {
				originLocation = location
				break
			}
		}

		rideStatus := model.RideStatus{
			RideLocationId:   originLocation.ID,
			RequestId:        request.ID,
			Status:           request.Status,
			CancellationNote: "",
		}
		err = tx.Create(&rideStatus).Error
		if err != nil {
			tx.Rollback()
			logger.Error("failed to create ride status", err)
			continue
		}

		err = storage.UnlockRequest(context.Background(), request.ID, lock)
		if err != nil {
			tx.Rollback()
			logger.Error("failed to release request lock", err, tag.Str("requestId", request.ID))
			continue
		}

		err = tx.Commit().Error
		if err != nil {
			logger.Error("failed to commit request", err, tag.Str("requestId", request.ID))
			continue
		}

		logger.Infof("starting scheduled request %v done successfully", request)
	}

	logger.Info("Create Request Job Scheduler Done Successfully")
}

func CalculateCourierPrice(customerId string, source *deliveryPb.CreateRequestLocation,
	destinations []*deliveryPb.CreateRequestLocation,
	vehicleType commonPb.VehicleType,
	requiredWorkers int32) (*pricingPb.CalculateCourierPriceResponse, error) {
	logger.Debugf("CalculateCourierPrice customerId %v source %v destinations %v vehicleType %v requiredWorkers %v", customerId, source, destinations, vehicleType, requiredWorkers)

	destinations2 := make([]*commonPb.Location, len(destinations))
	for i, d := range destinations {
		destinations2[i] = &commonPb.Location{
			Lat: d.Lat,
			Lon: d.Lon,
		}
	}
	query := &pricingPb.CalculateCourierPriceRequest{
		VehicleType:     vehicleType,
		RequiredWorkers: requiredWorkers,
		Source: &commonPb.Location{
			Lat: source.Lat,
			Lon: source.Lon,
		},
		Destinations: destinations2,
	}
	conn := getPricingConn()
	defer conn.Close()
	clientDeadline := time.Now().Add(time.Duration(60) * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	c := pricingPb.NewPricingClient(conn)

	logger.Debugf("CalculateCourierPrice response = %%v", c)

	return c.CalculateCourierPrice(ctx, query)
}

func getPricingConn() *grpc.ClientConn {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	opts = append(opts, grpc.WithBlock())
	logger.Infof("try to connect to pricing service %v", config.GetData().Pricing.GrpcPort)
	conn, err := grpc.Dial(config.GetData().Pricing.Connection, opts...)
	if err != nil {
		logger.Errorf("cannot connect to pricing: %v", err)
	}
	return conn
}
