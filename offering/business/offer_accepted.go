package business

import (
	"context"
	"fmt"
	"time"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	commonPb "github.com/kkjhamb01/courier-management/grpc/common/go"
	offeringPb "github.com/kkjhamb01/courier-management/grpc/offering/go"
	"github.com/kkjhamb01/courier-management/offering/db"
	"github.com/kkjhamb01/courier-management/offering/model"
	"github.com/kkjhamb01/courier-management/offering/pubsub"
	"github.com/kkjhamb01/courier-management/offering/storage"
)

func OnOfferAccepted(ctx context.Context, event offeringPb.OfferAcceptedEvent) (returnErr error) {
	logger.Infof("OnOfferAccepted event: %+v", event)
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

	isCourierValidToAccept, err := tx.IsCourierPendingOffer(ctx, event.OfferId, event.CourierId)
	if err != nil {
		logger.Error("failed to check if the courier is pending to the offer to check courier's eligibility to accept the offer", err)
		return err
	}
	if !isCourierValidToAccept {
		err = fmt.Errorf("the courier:%v is not related to the offer:%v (any more)", event.CourierId, event.OfferId)
		return err
	}

	// get the list of couriers assigned to the accepted OfferId
	courierIdsToCancel, err := tx.GetOfferPendingCouriers(ctx, event.OfferId)
	if err != nil {
		logger.Error("failed to get offer's pending couriers", err, tag.Str("offerId", event.OfferId))
		return err
	}
	// exclude the accepting courier from the list of cancelling couriers
	for i, courierIdToCancel := range courierIdsToCancel {
		if courierIdToCancel == event.CourierId {
			courierIdsToCancel = append(courierIdsToCancel[:i], courierIdsToCancel[i+1:]...)
			break
		}
	}

	// change the accepting courier status to ON_RIDE
	if err := tx.ChangeCourierStatus(ctx, event.CourierId, commonPb.CourierStatus_ON_RIDE); err != nil {
		logger.Error("failed to change courier status", err, tag.Str("courier id", event.CourierId))
		return err
	}

	// remove all couriers from the offer
	if err := tx.RemoveOfferAndAllPendingCouriers(ctx, event.OfferId); err != nil {
		logger.Error("failed to remove all pending couriers from the offer", err, tag.Str("offerId", event.OfferId))
		return err
	}

	if err := tx.CloseOffer(ctx, event.OfferId); err != nil {
		logger.Error("failed to close the offer", err, tag.Str("offerId", event.OfferId))
		return err
	}

	// TODO a better approach than iteration over all cancelling couriers ?
	// remove the offer from all cancelling couriers and the accepting courier
	for _, courierIdToCancel := range courierIdsToCancel {
		if err := tx.RemovePendingOfferFromCourier(ctx, event.OfferId, courierIdToCancel); err != nil {
			logger.Error("failed to remove the pending offer from the cancelling courier", err, tag.Str("offerId", event.OfferId), tag.Str("courier id", courierIdToCancel))
			return err
		}
	}
	if err := tx.RemovePendingOfferFromCourier(ctx, event.OfferId, event.CourierId); err != nil {
		logger.Error("failed to remove the pending offer from the accepting courier", err, tag.Str("offerId", event.OfferId), tag.Str("courier id", event.CourierId))
		return err
	}

	// publish the cancel event to all cancelling couriers
	for _, courierIdToCancel := range courierIdsToCancel {
		cancelEvent := offeringPb.OfferCancelledEvent{
			OfferId:      event.OfferId,
			CustomerId:   event.CustomerId,
			UpdatedBy:    event.CourierId,
			CancelledBy:  offeringPb.OfferCancelledEvent_COURIER,
			CancelReason: offeringPb.OfferCancelledEvent_ACCEPTED_BY_ANOTHER_COURIER,
		}

		// notify couriers
		pubsub.PublishById(pubsub.CancelOfferEvent(cancelEvent), courierIdToCancel)

		// publish on messaging system
		err = publishOfferCancelledEventOnMessaging(ctx, cancelEvent)
		if err != nil {
			logger.Error("failed to publish on messaging system", err)
			return err
		}
	}

	acceptedOfferModel := model.AcceptedOffer{
		CourierId:  event.CourierId,
		CustomerId: event.CustomerId,
		OfferId:    event.OfferId,
		Time:       time.Now(),
	}
	err = db.MariaDbClient().Create(&acceptedOfferModel).Error
	if err != nil {
		logger.Error("failed to insert accepted offer into database", err)
		return err
	}

	logger.Infof("offer accepted event handled successfully")
	return nil
}
