package business

import (
	"context"
	"fmt"
	"time"

	"github.com/kkjhamb01/courier-management/common/logger"
	commonPb "github.com/kkjhamb01/courier-management/grpc/common/go"
	offeringPb "github.com/kkjhamb01/courier-management/grpc/offering/go"
	"github.com/kkjhamb01/courier-management/offering/maps"
	"google.golang.org/protobuf/types/known/durationpb"
	googlemap "googlemaps.github.io/maps"
)

func bestEstimateArrivalTime(ctx context.Context, numberOfTopEstimates int, origins []commonPb.CourierLocation, dest offeringPb.Location) ([]commonPb.CourierETA, error) {

	logger.Infof("bestEstimateArrivalTime numberOfTopEstimates = %v, origins = %+v", numberOfTopEstimates, origins)

	courierETAs := make([]commonPb.CourierETA, 0, len(origins))

	// TODO: undo the changes to bring best estimate arrival back to the game
	for _, origin := range origins {
		courierETAs = append(courierETAs, commonPb.CourierETA{
			Courier:  origin.Courier,
			Duration: durationpb.New(time.Second * 20),
		})
	}

	//originsByVehicleTypes := make(map[commonPb.VehicleType][]commonPb.CourierLocation, 0)
	//for _, origin := range origins {
	//	if originsByVehicleTypes[origin.Courier.VehicleType] == nil {
	//		originsByVehicleTypes[origin.Courier.VehicleType] = make([]commonPb.CourierLocation, 0)
	//	}
	//
	//	originsByVehicleTypes[origin.Courier.VehicleType] = append(originsByVehicleTypes[origin.Courier.VehicleType], origin)
	//}
	//
	//for vehicleType, vehicleTypeOrigins := range originsByVehicleTypes {
	//	etasByVehicleTypes, err := bestEstimateArrivalTimeByVehicleType(ctx, vehicleTypeOrigins, dest, vehicleType)
	//	if err != nil {
	//		logger.Error("failed to get courier ETA", err)
	//		return nil, err
	//	}
	//	courierETAs = append(courierETAs, etasByVehicleTypes...)
	//}
	//
	//for _, newCourierETA := range courierETAs {
	//	var eta time.Duration
	//	indexToInsert := searchCourierETA(courierETAs, eta)
	//	courierETAs = append(courierETAs, commonPb.CourierETA{})
	//	copy(courierETAs[indexToInsert+1:], courierETAs[indexToInsert:])
	//	courierETAs[indexToInsert] = newCourierETA
	//}

	var lastIndex = numberOfTopEstimates
	if lastIndex > len(courierETAs) {
		lastIndex = len(courierETAs)
	}

	return courierETAs[:lastIndex], nil
}

func bestEstimateArrivalTimeByVehicleType(ctx context.Context, origins []commonPb.CourierLocation, dest offeringPb.Location, vehicleTime commonPb.VehicleType) ([]commonPb.CourierETA, error) {
	originsString := make([]string, len(origins))
	var i int
	for _, origin := range origins {
		originsString[i] = fmt.Sprint(origin.Location.Lat) + "," + fmt.Sprint(origin.Location.Lon)
		i++
	}

	var mode googlemap.Mode
	switch vehicleTime {
	case commonPb.VehicleType_BICYCLE:
		mode = googlemap.TravelModeBicycling
	default:
		mode = googlemap.TravelModeDriving
	}

	res, err := maps.GoogleClient().DistanceMatrix(ctx, &googlemap.DistanceMatrixRequest{
		Origins: originsString,
		Destinations: []string{
			fmt.Sprint(dest.Lat) + "," + fmt.Sprint(dest.Lon),
		},
		Mode:          mode,
		DepartureTime: "now",
		//`TrafficModelBestGuess`, `TrafficModelOptimistic`` or `TrafficModelPessimistic`.
		//TrafficModel:             "",
	})
	if err != nil {
		logger.Error("failed to get distance matrix from google map api", err)
		return nil, err
	}

	courierETAs := make([]commonPb.CourierETA, 0, len(origins))
CourierLoop:
	for index, resRow := range res.Rows {
		var eta time.Duration
		var distance int
		for _, element := range resRow.Elements {
			if element.Status != maps.GoogleStatusOK {
				continue CourierLoop
			}
			eta += element.DurationInTraffic
			distance += element.Distance.Meters
		}

		newCourierETA := commonPb.CourierETA{
			Courier: &commonPb.Courier{
				Id:          origins[index].Courier.Id,
				VehicleType: origins[index].Courier.VehicleType,
			},
			Duration: durationpb.New(eta),
			Meters:   int32(distance),
		}

		courierETAs = append(courierETAs, newCourierETA)
	}

	return courierETAs, nil
}

func searchCourierETA(etas []commonPb.CourierETA, key time.Duration) int {
	if len(etas) == 0 {
		return 0
	}

	left := 0
	right := len(etas)
	center := (left + right) / 2

	for left < right {
		if etas[center].Duration.AsDuration() == key {
			return center
		} else if etas[center].Duration.AsDuration() > key {
			right = center - 1
			center = (left + right) / 2
		} else {
			left = center + 1
			center = (left + right) / 2
		}
	}

	if left >= len(etas) {
		return len(etas)
	}

	if int64(etas[left].Duration.Nanos) > key.Nanoseconds() {
		return left
	} else {
		return left + 1
	}
}
