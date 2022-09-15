package business

import (
	"context"
	"errors"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/offering/db"
	"gitlab.artin.ai/backend/courier-management/offering/model"
	"gorm.io/gorm"
)

func HadCustomerRideWithCourier(ctx context.Context, customerId string, courierId string, offerId string) (bool, error) {

	logger.Infof("HadCustomerRideWithCourier courierId = %v, customerId: %v, offerId = %v", courierId, customerId, offerId)

	var acceptedOffer model.AcceptedOffer
	err := db.MariaDbClient().
		First(&acceptedOffer, "customer_id = ? AND courier_id = ? AND offer_id = ?", customerId, courierId, offerId).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		logger.Error("failed to fetch acceptedOffer", err)
		return false, err
	}

	return acceptedOffer.ID != "", nil
}

func GetOfferCustomerAndCourier(ctx context.Context, offerId string) (customerId string, courierId string, _err error) {

	logger.Infof("GetOfferCustomerAndCourier offerId = %v", offerId)

	var acceptedOffer model.AcceptedOffer
	err := db.MariaDbClient().
		First(&acceptedOffer, "offer_id = ?", offerId).
		Error

	if err != nil {
		logger.Error("failed to fetch acceptedOffer", err)
		return "", "", err
	}

	return acceptedOffer.CustomerId, acceptedOffer.CourierId, nil
}
