package business

import (
	"context"
	"fmt"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/kkjhamb01/courier-management/offering/storage"
)

func OnOfferRejected(ctx context.Context, offerId string, courierId string) (returnErr error) {

	logger.Infof("OnOfferRejected courierId = %v, offerId = %v", courierId, offerId)

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

	isCourierValidToReject, err := tx.IsCourierPendingOffer(ctx, offerId, courierId)
	if err != nil {
		logger.Error("failed to check if the courier is pending to the offer to check courier's eligibility to reject the offer", err)
		return err
	}
	if !isCourierValidToReject {
		err = fmt.Errorf("the courier:%v is not related to the offer:%v (any more)", courierId, offerId)
		return err
	}

	if err := tx.RemovePendingOfferFromCourier(ctx, offerId, courierId); err != nil {
		logger.Error("failed to remove offer from the courier", err, tag.Str("courierId", courierId), tag.Str("offerId", offerId))
		return err
	}

	if err := tx.RemovePendingCourierFromOffer(ctx, offerId, courierId); err != nil {
		logger.Error("failed to remove courier from the offer", err, tag.Str("courierId", courierId), tag.Str("offerId", offerId))
		return err
	}

	logger.Info("offer rejected event handled successfully")
	return nil
}
