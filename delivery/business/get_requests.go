package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/delivery/db"
	"gitlab.artin.ai/backend/courier-management/delivery/model"
	deliveryPb "gitlab.artin.ai/backend/courier-management/grpc/delivery/go"
	uaa "gitlab.artin.ai/backend/courier-management/uaa/proto"
)

func GetRequests(ctx context.Context, req *deliveryPb.GetRequestsRequest) ([]*deliveryPb.Request, error) {

	logger.Infof("GetRequests request = %+v", req)

	var requestModels []model.Request
	result := db.MariaDbClient().Preload("Locations")

	if len(req.GetId()) > 0 {
		result = result.Where("ID = ?", req.GetId()).First(&requestModels)
	} else if len(req.GetCustomerId()) > 0 {
		result = result.Where("CUSTOMER_ID = ?", req.GetCustomerId()).Find(&requestModels)
	} else if len(req.GetCourierId()) > 0 {
		result = result.Where("COURIER_ID = ?", req.GetCourierId()).Find(&requestModels)
	} else if req.GetDestination() != nil {
		result = result.
			Joins("left join request_locations on requests.id = request_locations.request_id").
			Where("POWER(request_locations.LAT - ?, 2) + POWER(request_locations.LON - ?, 2) < ? AND "+
				"request_locations.is_origin = 0",
				req.GetDestination().Lat, req.GetDestination().Lon, req.GetDestination().Range).
			Find(&requestModels)
	} else if req.GetOrigin() != nil {
		result = result.
			Joins("left join request_locations on requests.id = request_locations.request_id").
			Where("POWER(request_locations.LAT - ?, 2) + POWER(request_locations.LON - ?, 2) < ? AND "+
				"request_locations.is_origin = 1",
				req.GetOrigin().Lat, req.GetOrigin().Lon, req.GetOrigin().Range).
			Find(&requestModels)
	} else if req.GetPrice() != nil {
		priceValue := req.GetPrice().Units + (int64(req.GetPrice().Nanos) / 1_000_000_000)
		result = result.Where("FINAL_PRICE = ? AND FINAL_PRICE_CURRENCY = ?", priceValue, req.GetPrice().CurrencyCode).Find(&requestModels)
	} else if req.GetTimeRange() != nil {
		result = result.Where("CREATED_AT >= ? AND CREATED_AT <= ? ",
			req.GetTimeRange().From.AsTime(), req.GetTimeRange().To.AsTime()).Find(&requestModels)
	}

	if result.Error != nil {
		logger.Error("failed fetch request", result.Error)
		return nil, uaa.Internal.Error(result.Error)
	}

	requestProtos := make([]*deliveryPb.Request, len(requestModels))
	for i, requestModel := range requestModels {
		requestProtos[i] = requestModel.ToProtoP()
	}

	return requestProtos, nil
}

func GetCourierActiveRequest(ctx context.Context, courierId string) (*deliveryPb.GetCourierActiveRequestResponse, error) {

	logger.Infof("GetCourierActiveRequest courierId = %v", courierId)

	var requestModels model.Request
	err := db.MariaDbClient().
		Preload("Locations").
		Where("courier_id = ?", courierId).
		Where("status <> ?", deliveryPb.RequestStatus_COMPLETED.String()).
		Where("status <> ?", deliveryPb.RequestStatus_COURIER_CANCELLED.String()).
		Where("status <> ?", deliveryPb.RequestStatus_CUSTOMER_CANCELLED.String()).
		Order("created_at DESC").
		First(&requestModels).Error
	if err != nil {
		logger.Error("failed to fetch requests", err)
		return nil, uaa.Internal.Error(err)
	}

	var requestStatuses []model.RideStatus
	err = db.MariaDbClient().
		Where("request_id = ?", requestModels.ID).
		Order("created_at asc").
		Find(&requestStatuses).
		Error
	if err != nil {
		logger.Error("failed to fetch request statuses", err)
		return nil, uaa.Internal.Error(err)
	}

	requestStatusesProto := make([]*deliveryPb.RideStatus, len(requestStatuses))
	for i, requestStatus := range requestStatuses {
		requestStatusesProto[i] = requestStatus.ToProtoP()
	}

	customerProfile, err := getCustomerProfile(ctx, requestModels.CustomerId)
	if err != nil {
		logger.Error("failed to fetch request statuses", err)
		return nil, uaa.Internal.Error(err)
	}

	requestHistoryItem := &deliveryPb.RequestHistoryItem{
		Request: requestModels.ToProtoP(),
		Events:  requestStatusesProto,
	}

	return &deliveryPb.GetCourierActiveRequestResponse{
		RequestAndStatuses: requestHistoryItem,
		RequesterName:      customerProfile.FirstName + " " + customerProfile.LastName,
		RequesterPhone:     customerProfile.PhoneNumber,
	}, nil
}

func GetCustomerActiveRequest(ctx context.Context, customerId string) (*deliveryPb.GetCustomerActiveRequestResponse, error) {

	logger.Infof("GetCustomerActiveRequest customerId = %v", customerId)

	var requestModels model.Request
	err := db.MariaDbClient().
		Preload("Locations").
		Where("customer_id = ?", customerId).
		Where("status <> ?", deliveryPb.RequestStatus_COMPLETED.String()).
		Where("status <> ?", deliveryPb.RequestStatus_COURIER_CANCELLED.String()).
		Where("status <> ?", deliveryPb.RequestStatus_CUSTOMER_CANCELLED.String()).
		Order("created_at DESC").
		First(&requestModels).Error
	if err != nil {
		logger.Error("failed to fetch requests", err)
		return nil, uaa.Internal.Error(err)
	}

	var requestStatuses []model.RideStatus
	err = db.MariaDbClient().
		Where("request_id = ?", requestModels.ID).
		Order("created_at asc").
		Find(&requestStatuses).
		Error
	if err != nil {
		logger.Error("failed to fetch request statuses", err)
		return nil, uaa.Internal.Error(err)
	}

	requestStatusesProto := make([]*deliveryPb.RideStatus, len(requestStatuses))
	for i, requestStatus := range requestStatuses {
		requestStatusesProto[i] = requestStatus.ToProtoP()
	}

	customerProfile, err := getCustomerProfile(ctx, requestModels.CustomerId)
	if err != nil {
		logger.Error("failed to fetch request statuses", err)
		return nil, uaa.Internal.Error(err)
	}

	requestHistoryItem := &deliveryPb.RequestHistoryItem{
		Request: requestModels.ToProtoP(),
		Events:  requestStatusesProto,
	}

	return &deliveryPb.GetCustomerActiveRequestResponse{
		RequestAndStatuses: requestHistoryItem,
		RequesterName:      customerProfile.FirstName + " " + customerProfile.LastName,
		RequesterPhone:     customerProfile.PhoneNumber,
	}, nil
}
