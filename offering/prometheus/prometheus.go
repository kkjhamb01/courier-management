package prometheus

import (
	"context"
	"time"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/prometheus"
	"github.com/kkjhamb01/courier-management/offering/storage"
	prometheusClient "github.com/prometheus/client_golang/prometheus"
)

func Setup(ctx context.Context) {
	prometheus.AddGauge("offering_available_couriers",
		"The total number of available couriers", func(gauge prometheusClient.Gauge) {
			var numberOfAvailableCouriers float64
			// TODO fetch from storage
			gauge.Set(numberOfAvailableCouriers)
		}, 1*time.Minute)

	prometheus.AddGauge("offering_unavailable_couriers",
		"The total number of unavailable couriers", func(gauge prometheusClient.Gauge) {
			var numberOfUnavailableCouriers float64
			// TODO fetch from storage
			gauge.Set(numberOfUnavailableCouriers)
		}, 1*time.Minute)

	prometheus.AddGauge("offering_blocked_couriers",
		"The total number of blocked couriers", func(gauge prometheusClient.Gauge) {
			var numberOfBlockedCouriers float64
			// TODO fetch from storage
			gauge.Set(numberOfBlockedCouriers)
		}, 1*time.Minute)

	prometheus.AddGauge("offering_on_ride_couriers",
		"The total number of on ride couriers", func(gauge prometheusClient.Gauge) {
			var numberOfOnRideCouriers float64
			// TODO fetch from storage
			gauge.Set(numberOfOnRideCouriers)
		}, 1*time.Minute)

	prometheus.AddGauge("offering_pending_offers",
		"The total number of offers pending for an answer from a courier", func(gauge prometheusClient.Gauge) {
			var numberOfPendingOffers int
			var err error
			tx := storage.CreateTx()
			defer tx.Rollback()

			numberOfPendingOffers, err = tx.NumberOfPendingOffers(ctx)
			if err != nil {
				logger.Error("failed to get number of pending offers", err)
				return
			}

			gauge.Set(float64(numberOfPendingOffers))
		}, 1*time.Minute)
}
