package service

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
)

const (
	Finance      = "finance"
	Offering     = "offering"
	Pricing      = "pricing"
	Party        = "party"
	Uaa          = "uaa"
	Notification = "notification"
	Delivery     = "delivery"
	Rating       = "rating"
	Promotion    = "promotion"
	Announcement    = "announcement"

	Prometheus = "prometheus"
)

func AllAvailableServices() []string {
	return []string{
		Finance,
		Offering,
		Pricing,
		Party,
		Uaa,
		Notification,
		Delivery,
		Rating,
		Promotion,
		Announcement,
		Prometheus,
	}
}

type Service interface {
	Start(context.Context)
	Stop() error
	Name() string
	// the relative path to the service directory from the project root
	RelDir() string
}

var RegisteredServices []Service

func RegisterService(newService Service) {
	RegisteredServices = append(RegisteredServices, newService)
}

func FilterRegisteredServices(filter func(s Service) bool) {
	var tmp []Service
	for _, registeredService := range RegisteredServices {
		if !filter(registeredService) {
			tmp = append(tmp, registeredService)
		}
	}
	RegisteredServices = tmp
}

func StartAllRegisteredServices(ctx context.Context) {
	// TODO Consider concurrency here (Important)
	for _, rs := range RegisteredServices {
		rs.Start(ctx)
	}
}

func StopAllRegisteredServices() {
	for _, rs := range RegisteredServices {
		err := rs.Stop()
		if err != nil {
			logger.Error("failed to stop the service", err, tag.Str("service name", rs.Name()))
		}
		logger.Info("service stopped", tag.Str("service name", rs.Name()))
	}
}
