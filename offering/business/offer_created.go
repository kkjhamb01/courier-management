package business

import (
	"context"
	"errors"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	commonPb "gitlab.artin.ai/backend/courier-management/grpc/common/go"
	offeringPb "gitlab.artin.ai/backend/courier-management/grpc/offering/go"
	"gitlab.artin.ai/backend/courier-management/offering/pubsub"
	"gitlab.artin.ai/backend/courier-management/offering/storage"
	"google.golang.org/protobuf/types/known/durationpb"
	"time"
)

func OnNewOfferCreated(ctx context.Context, offer offeringPb.Offer) (returnErr error) {
	logger.Infof("OnNewOfferCreated offer: %+v", offer)
	tx := storage.CreateTx()

	defer func() {
		if returnErr != nil {
			rollbackErr := tx.Rollback()
			logger.Error("failed to rollback tx", rollbackErr)
			return
		}

		returnErr = tx.Commit(ctx)
		if returnErr != nil {
			logger.Error("failed to commit tx", returnErr)
		}
	}()

	// get already sent courierLocations
	sentCouriers, err := tx.GetOfferSentCouriers(ctx, offer.Id)
	if err != nil {
		logger.Error("failed to get sent courierLocations", err, tag.Obj("offer id", offer.Id))
		return err
	}
	logger.Infof("OnNewOfferCreated: GetOfferSentCouriers %v", sentCouriers)

	// get a number of courierLocations near the offer source
	courierLocations, err := tx.GetNearbyCouriersByCourierTypeAndMaxResult(ctx,
		commonPb.Location{
			Lat: offer.Source.Lat,
			Lon: offer.Source.Lon,
		},
		config.Offering().CouriersDistanceOnNewOffer,
		offer.VehicleType,
		config.Offering().NearbyCouriersToSearch)
	if err != nil {
		logger.Error("failed to search courierLocations nearby on new offer", err, tag.Obj("offer", offer))
		return err
	}
	// exclude previously sent couriers
	for i := len(courierLocations) - 1; i >= 0; i-- {
		courierLocation := courierLocations[i]
		for _, sentCourier := range sentCouriers {
			if sentCourier == courierLocation.Courier.Id {
				courierLocations = append(courierLocations[:i], courierLocations[i+1:]...)
				break
			}
		}
	}
	logger.Infof("OnNewOfferCreated: nearby courierLocations: %v", courierLocations)

	for i := len(courierLocations) - 1; i >= 0; i-- {
		courierLocation := courierLocations[i]
		isActive, err := IsCourierActive(ctx, courierLocation.Courier.Id)
		if err != nil {
			logger.Error("failed to check if the courierLocation is active", err)
			return err
		}
		if !isActive {
			logger.Infof("the courierLocation %v is not active", courierLocation.Courier.Id)
			courierLocations = append(courierLocations[:i], courierLocations[i+1:]...)
		}
	}

	courierETAs, err := bestEstimateArrivalTime(ctx, config.Offering().MaxCouriersOnNewOffer, courierLocations, *offer.Source)
	if err != nil {
		logger.Error("failed to get best estimate arrival times", err)
		return err
	}
	logger.Infof("OnNewOfferCreated: bestEstimateArrivalTime: %v", courierETAs)

	if len(courierETAs) == 0 {
		// TODO specify a policy on keep looking
		err = errors.New("no courierLocation found near the offer source")
		logger.Error("failed to find courierLocations near the source", err, tag.Obj("offer", offer))
		return err
	}

	// get courierLocations status
	couriersStatus, err := tx.GetCouriersStatusByCourierETAs(ctx, courierETAs)
	if err != nil {
		logger.Error("failed to check courierLocations status", err, tag.Obj("courierLocations", courierETAs))
		return err
	}

	// filter courierLocations
	for i := 0; i < len(courierETAs); i++ {
		// check the courierLocation status to filter unavailable drivers
		isAvailable := couriersStatus[i] == commonPb.CourierStatus_AVAILABLE

		// TODO performance bottleneck: get all courierLocations pending offers at once (?)
		// check the courierLocation pending offers to prevent assigning more than the max pendingOffers to a courierLocations
		pendingOffers, err := tx.GetCourierPendingOffers(ctx, courierETAs[i].Courier.Id)
		if err != nil {
			logger.Error("failed to get courierLocation pending offers", err, tag.Str("courierLocation id", courierETAs[i].Courier.Id))
			return err
		}
		isExceedingMaxOffers := len(pendingOffers) >= config.Offering().MaxOffersPerCourier && config.Offering().MaxOffersPerCourier > 0

		// remove the courierLocation from the list if :
		// 		the courierLocation is unavailable
		//		OR
		// 		the courierLocation has been offered for the maximum number of permitted times
		if !isAvailable || isExceedingMaxOffers {
			courierETAs = append(courierETAs[:i], courierETAs[i+1:]...)
			couriersStatus = append(couriersStatus[0:i], couriersStatus[i+1:]...)
			i--
		}
	}
	// define slice of courier ids
	courierIds := make([]string, len(courierETAs), len(courierETAs))
	for i, courierLocation := range courierETAs {
		courierIds[i] = courierLocation.Courier.Id
	}

	// add pending courierLocations to the offer
	err = tx.AddPendingCouriersToOffer(ctx, offer.Id, courierIds)
	if err != nil {
		logger.Error("failed to assign pending courierLocations to the offer", err,
			tag.Str("offer id", offer.Id), tag.Obj("courierIds", courierIds))
		return err
	}

	// add sent courierLocations to the offer
	err = tx.AddOfferSentCourier(ctx, offer.Id, courierIds)
	if err != nil {
		logger.Error("failed to assign pending courierLocations to the offer", err,
			tag.Str("offer id", offer.Id), tag.Obj("courierIds", courierIds))
		return err
	}

	// add pending offer to the courierLocations
	// TODO Performance bottleneck: assign the offer to all courierLocations at once (?)
	for _, courierId := range courierIds {
		err = tx.AddPendingOfferToCourier(ctx, offer.Id, courierId)
		if err != nil {
			logger.Error("failed to add the courierLocation to the offer", err,
				tag.Str("offer id", offer.Id), tag.Obj("courierId", courierId))
			return err
		}
	}

	// publish new offer to subscribed courierLocations
	// TODO Performance bottleneck: publish single event with list of courierLocations vs send multiple events
	var customerName string
	var customerPhone string
	customerProfile, err := getCustomerProfile(ctx, offer.CustomerId)
	if err != nil {
		logger.Error("failed to fetch customer info from party", err)
		customerName = "failed to fetch"
	} else {
		customerName = customerProfile.FirstName + " " + customerProfile.LastName
		customerPhone = customerProfile.PhoneNumber
	}
	for index, courierId := range courierIds {
		pubsub.PublishById(pubsub.NewOfferEvent(offeringPb.NewOfferEvent{
			Offer:                  &offer,
			CourierId:              courierId,
			CourierResponseTimeout: durationpb.New(config.Offering().CourierTimeToAnswerOffer),
			RequesterName:          customerName,
			RequesterPhone:         customerPhone,
			Desc:                   "new offer is available",
			Duration:               courierETAs[index].Duration,
			DistanceMeters:         courierETAs[index].Meters,
		}), courierId)
	}
	//
	//
	// publish on messaging
	newOfferSentEvent := offeringPb.NewOfferSentToCouriersEvent{
		Offer:      &offer,
		CourierIds: courierIds,
	}
	err = publishNewOfferSentToCouriersEventOnMessaging(ctx, newOfferSentEvent)
	if err != nil {
		logger.Error("failed to publish on messaging system", err)
		return err
	}

	go func(ctx context.Context, offer offeringPb.Offer) {
		select {
		case <-time.After(config.Offering().CourierTimeToAnswerOffer):
			logger.Infof("creating offer: timeout for offer id : %v", offer.Id)
			err := OnOfferTimeout(ctx, offer)
			if err != nil {
				logger.Error("failed to perform expected action after offer expiration", err)
			}
			return
		}
	}(ctx, offer)

	logger.Infof("new created offer event handled successfully")
	return nil
}
