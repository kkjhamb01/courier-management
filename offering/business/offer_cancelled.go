package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	offeringPb "gitlab.artin.ai/backend/courier-management/grpc/offering/go"
	"gitlab.artin.ai/backend/courier-management/offering/pubsub"
	"gitlab.artin.ai/backend/courier-management/offering/storage"
)

func OnOfferCancelled(ctx context.Context, event offeringPb.OfferCancelledEvent) (returnErr error) {
	logger.Infof("OnOfferCancelled event: %+v", event)
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

	// get the list of couriers assigned to the accepted OfferId
	courierIdsToCancel, err := tx.GetOfferPendingCouriers(ctx, event.OfferId)
	if err != nil {
		logger.Error("failed to get offer's pending couriers", err, tag.Str("offerId", event.OfferId))
		return err
	}

	// remove all pending couriers from the offer
	if err := tx.RemoveOfferAndAllPendingCouriers(ctx, event.OfferId); err != nil {
		logger.Error("failed to remove offer from the courier", err, tag.Str("offerId", event.OfferId))
		return err
	}

	// remove pending offer from all cancelling couriers
	for _, courierIdToCancel := range courierIdsToCancel {
		if err := tx.RemovePendingOfferFromCourier(ctx, event.OfferId, courierIdToCancel); err != nil {
			logger.Error("failed to remove courier from the offer", err, tag.Str("courierId", courierIdToCancel), tag.Str("offerId", event.OfferId))
			return err
		}
	}

	// publish cancelling offer to pending couriers and messaging system
	for _, courierIdToCancel := range courierIdsToCancel {
		// notify couriers
		pubsub.PublishById(pubsub.CancelOfferEvent(event), courierIdToCancel)
	}

	if event.CancelReason != offeringPb.OfferCancelledEvent_TIMEOUT &&
		event.CancelReason != offeringPb.OfferCancelledEvent_ACCEPTED_BY_ANOTHER_COURIER {
		if err = tx.CloseOffer(ctx, event.OfferId); err != nil {
			logger.Error("failed to close the offer", err, tag.Str("offerId", event.OfferId))
			return err
		}
	}

	logger.Infof("cancelled offer event handled successfully")
	return nil
}
