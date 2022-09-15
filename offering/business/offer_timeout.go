package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	offeringPb "gitlab.artin.ai/backend/courier-management/grpc/offering/go"
	"gitlab.artin.ai/backend/courier-management/offering/pubsub"
	"gitlab.artin.ai/backend/courier-management/offering/storage"
)

func OnOfferTimeout(ctx context.Context, offer offeringPb.Offer) (returnErr error) {

	logger.Infof("OnOfferTimeout offer = %+v", offer)

	logger.Info("on offer timeout")

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

	// return if the offer is closed (done)
	isOfferClosed, err := tx.IsOfferClosed(ctx, offer.Id)
	if err != nil {
		logger.Error("failed to check if the offer is closed", err)
		return err
	}
	if isOfferClosed {
		logger.Infof("no need to cancel anything on timeout")
		return nil
	}

	// get the list of couriers assigned to the cancelling OfferId
	courierIdsToCancel, err := tx.GetOfferPendingCouriers(ctx, offer.Id)
	if err != nil {
		logger.Error("failed to get offer's pending couriers", err, tag.Str("offerId", offer.Id))
		return err
	}

	// remove the offer and all couriers
	if err := tx.RemoveOfferAndAllPendingCouriers(ctx, offer.Id); err != nil {
		logger.Error("failed to remove the all pending couriers from the offer", err, tag.Str("offerId", offer.Id))
		return err
	}

	// remove pending offer from cancelling courier
	for _, courierIdToCancel := range courierIdsToCancel {
		if err := tx.RemovePendingOfferFromCourier(ctx, offer.Id, courierIdToCancel); err != nil {
			logger.Error("failed to remove the pending offer from the courier", err, tag.Str("courierId", courierIdToCancel), tag.Str("offerId", offer.Id))
			return err
		}
	}

	// publish cancelling offer to pending couriers and messaging system
	for _, courierIdToCancel := range courierIdsToCancel {
		cancelEvent := offeringPb.OfferCancelledEvent{
			OfferId:      offer.Id,
			CustomerId:   offer.CustomerId,
			UpdatedBy:    "",
			CancelReason: offeringPb.OfferCancelledEvent_TIMEOUT,
			CancelledBy:  offeringPb.OfferCancelledEvent_SYSTEM,
		}

		// notify couriers
		pubsub.PublishById(pubsub.CancelOfferEvent(cancelEvent), courierIdToCancel)

		// publish on messaging
		err = publishOfferCancelledEventOnMessaging(ctx, cancelEvent)
		if err != nil {
			logger.Error("failed to publish cancel event", err)
			return err
		}
	}

	offerRetries, err := tx.GetOfferRetries(ctx, offer.Id)
	if err != nil {
		logger.Error("failed to get offer retries", err)
		return err
	}

	// if offer retries has reached the max, publish its event and return
	if offerRetries >= config.Offering().MaxOfferRetries {
		// publish max retries reached on messaging
		maxRetiesEvent := offeringPb.MaxOfferRetriesEvent{
			OfferId:    offer.Id,
			CustomerId: offer.CustomerId,
			Desc:       "",
		}
		err = publishMaxRetriesOnMessaging(ctx, maxRetiesEvent)
		if err != nil {
			logger.Error("failed to publish on messaging system", err)
			return err
		}

		return nil
	}

	_, err = tx.IncreaseOfferRetries(ctx, offer.Id)
	if err != nil {
		logger.Error("failed to increase offer retry", err)
		return err
	}

	// request for offer retry
	// publish on messaging
	retryRequestEvent := offeringPb.RetryOfferRequestEvent{
		Offer: &offer,
		Desc:  "",
	}
	err = publishRetryOfferRequestedOnMessaging(ctx, retryRequestEvent)
	if err != nil {
		logger.Error("failed to publish on messaging system", err)
		return err
	}

	return nil
}
