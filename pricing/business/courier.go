package business

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	commonPb "gitlab.artin.ai/backend/courier-management/grpc/common/go"
	pricingPb "gitlab.artin.ai/backend/courier-management/grpc/pricing/go"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
)

func CalculateCourierPrice(ctx context.Context, req *pricingPb.CalculateCourierPriceRequest) (*pricingPb.CalculateCourierPriceResponse, error) {
	distance, duration, err := callDistanceApi(req.Source, req.Destinations)

	if err != nil {
		logger.Error("failed to call distance api", err)
		return nil, err
	}

	logger.Debugf("CalculateCourierPrice distance = %v, duration = %v", distance, duration)

	price := calculatePrice(distance, duration, req.GetVehicleType(), req.GetRequiredWorkers())

	res := &pricingPb.CalculateCourierPriceResponse{
		EstimatedDuration: duration,
		EstimatedDistance: distance,
		Amount: float64(price),
		Currency: "GBP",
	}
	logger.Debugf("CalculateCourierPrice result = %%v", res)
	return res, nil
}

func ReviewCourierPrice(ctx context.Context, req *pricingPb.ReviewCourierPriceRequest) (*pricingPb.ReviewCourierPriceResponse, error) {
	distance, duration, err := callDistanceApi(req.Source, req.Destinations)

	if err != nil {
		logger.Error("failed to call distance api", err)
		return nil, err
	}

	logger.Debugf("CalculateCourierPrice distance = %v, duration = %v", distance, duration)

	vehicleTypes := []commonPb.VehicleType{
		commonPb.VehicleType_BICYCLE,
		commonPb.VehicleType_SMALL_VAN,
		commonPb.VehicleType_MEDIUM_VAN,
		commonPb.VehicleType_LARGE_VAN,
		commonPb.VehicleType_CAR,
		commonPb.VehicleType_MOTORBIKE,
		commonPb.VehicleType_TRUCK,
	}

	var prices = make([]*pricingPb.ReviewCourierPriceResponse_Price, len(vehicleTypes))
	for i,vt := range vehicleTypes {
		prices[i] = &pricingPb.ReviewCourierPriceResponse_Price{
			VehicleType: vt,
			Currency: "GBP",
			Amount: float64(calculatePrice(distance, duration, vt, req.RequiredWorkers)),
		}
	}

	return &pricingPb.ReviewCourierPriceResponse{
		EstimatedDuration: duration,
		EstimatedDistance: distance,
		Prices: prices,
	}, nil
}

func calculatePrice(distance int32, duration int64, vehicleType commonPb.VehicleType, numberOfWorkers int32) int32{
	return 3 * distance / 1000 + 5 * int32(duration) / 60 + numberOfWorkers * 7
}

func callDistanceApi(source *commonPb.Location, destinations []*commonPb.Location) (int32, int64, error) {
	destinationsStr := make([]string, len(destinations))
	for i,d := range destinations {
		destinationsStr[i] = fmt.Sprintf("%v,%v", d.Lat, d.Lon)
	}
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/distancematrix/json?origins=%v&destinations=%v&key=%v",
		fmt.Sprintf("%v,%v", source.Lat, source.Lon),
		strings.Join(destinationsStr[:], "%7C"),
		config.GetData().DistanceApi.Key,
	)

	logger.Debugf("callDistanceApi url", url)

	method := "GET"

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return 0,0,err
	}
	res, err := client.Do(req)
	if err != nil {
		return 0,0,err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0,0,err
	}
	response := GoogleDistanceResponse{}
	logger.Debugf("callDistanceApi response body = %v", string(body))
	json.Unmarshal(body, &response)
	var distance int32
	var duration int64
	if response.Rows != nil && len(response.Rows) > 0 {
		row := response.Rows[0]
		if row.Elements != nil {
			for _,element := range row.Elements {
				if element.Distance == nil{
					return 0, 0, errors.New("empty response from Google")
				}
				distance = distance + element.Distance.Value
				duration = duration + element.Duration.Value
			}
		}
	}
	return distance, duration, nil
}

func roundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

type GoogleDistanceResponse struct {
	Rows     	 []*Row `json:"rows"`
}
type Row struct {
	Elements     []*Element `json:"elements"`
}
type Element struct {
	Distance     *Distance `json:"distance"`
	Duration     *Duration `json:"duration"`
}
type Distance struct {
	Value     int32 	`json:"value"`
}
type Duration struct {
	Value     int64 `json:"value"`
}