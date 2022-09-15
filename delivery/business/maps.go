package business

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	deliveryPb "github.com/kkjhamb01/courier-management/grpc/delivery/go"
)

func isOriginAndDestinationWithinAllowedCities(ctx context.Context, origin *deliveryPb.CreateRequestLocation, destinations []*deliveryPb.CreateRequestLocation) (bool, error) {

	logger.Infof("isOriginAndDestinationWithinAllowedCities origin = %+v, destinations = %+v", origin, destinations)

	ok, err := isWithinAllowedCities(ctx, origin.Lat, origin.Lon, config.Delivery().ValidCities)
	if err != nil {
		return false, err
	}
	if !ok {
		logger.Info("origin is not within valid range", tag.Obj("origin", origin), tag.Obj("valid cities", config.Delivery().ValidCities))
		return false, nil
	}

	for _, destination := range destinations {
		ok, err := isWithinAllowedCities(ctx, destination.Lat, destination.Lon, config.Delivery().ValidCities)
		if err != nil {
			return false, err
		}
		if !ok {
			logger.Info("destination is not within valid range", tag.Obj("destination", destination), tag.Obj("valid cities", config.Delivery().ValidCities))
			return false, nil
		}
	}

	return true, nil
}

func isWithinAllowedCities(ctx context.Context, lat float64, lon float64, allowedCities []string) (bool, error) {
	//	geoResults, err := maps.GoogleClient().ReverseGeocode(ctx, &googlemaps.GeocodingRequest{
	//		LatLng: &googlemaps.LatLng{
	//			Lat: lat,
	//			Lng: lon,
	//		},
	//	})
	//	if err != nil {
	//		logger.Error("failed to get google map reverse gecode result", err)
	//		return false, err
	//	}
	//
	//	var townNameAddressComponent *googlemaps.AddressComponent
	//resultLoop:
	//	for _, geoResult := range geoResults {
	//		for _, addressComponent := range geoResult.AddressComponents {
	//			for _, addressComponentType := range addressComponent.Types {
	//				if addressComponentType == "locality" {
	//					townNameAddressComponent = &addressComponent
	//					break resultLoop
	//				}
	//			}
	//
	//		}
	//	}
	//
	//	if townNameAddressComponent == nil {
	//		err = fmt.Errorf("no postal_town addressComponentType found on the address: lat/long: %v, %v", lat, lon)
	//		logger.Error("failed to fetch lat/lon city name", err)
	//		return false, err
	//	}
	//
	//	for _, allowedCity := range allowedCities {
	//		if allowedCity == townNameAddressComponent.LongName || allowedCity == townNameAddressComponent.ShortName {
	//			return true, nil
	//		}
	//	}

	return true, nil
}
