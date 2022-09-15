package business

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/party/domain"
	pb "gitlab.artin.ai/backend/courier-management/party/proto"
	"gitlab.artin.ai/backend/courier-management/uaa/proto"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type MotErrors struct {
	Errors []MotError `json:"errors"`
}

type MotError struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (s *Service) SearchMot(ctx context.Context, in *pb.SearchMotRequest) (*pb.SearchMotResponse, error) {
	logger.Infof("SearchMot registrationNumber = %v", in.GetRegistrationNumber())

	values := map[string]string{"registrationNumber": in.GetRegistrationNumber()}

	jsonValue, _ := json.Marshal(values)

	req, err := http.NewRequest("POST", s.config.Mot.ApiAddress, bytes.NewBuffer(jsonValue))
	req.Header.Set("x-api-key", s.config.Mot.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Duration(s.config.Mot.Timeout) * time.Second,
	}
	resp, err := client.Do(req)

	if err != nil {
		logger.Errorf("SearchMot cannot send request to MOT", err)
		return nil, proto.Internal.Error(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	switch resp.StatusCode {
	case 400:
		var motError MotErrors
		json.Unmarshal(body, &motError)
		logger.Debugf("SearchMot mot bad request for registration number %v ; error = %v", in.GetRegistrationNumber(), motError)
		var details = make([]string, len(motError.Errors))
		for i, motError := range motError.Errors {
			details[i] = motError.Detail
		}
		detailsStr := strings.Join(details[:], ",")
		return nil, proto.InvalidArgument.ErrorMsg(detailsStr)
	case 404:
		logger.Debugf("SearchMot mot not found %v", in.GetRegistrationNumber())
		return nil, proto.NotFound.ErrorMsg(resp.Status)
	case 500, 503:
		logger.Debugf("SearchMot mot error from server %v", in.GetRegistrationNumber())
		return nil, proto.Unavailable.ErrorMsg(resp.Status)
	case 401, 403:
		logger.Debugf("SearchMot request to mot server is unauthorized")
		return nil, proto.Unauthenticated.ErrorNoMsg()
	case 200:
		logger.Debugf("SearchMot response from mot %v", string(body))
		var courierMot domain.CourierMot
		err = json.Unmarshal(body, &courierMot)
		if err != nil {
			logger.Errorf("SearchMot cannot decode MOT", err)
			return nil, proto.Internal.Error(err)
		}
		logger.Debugf("SearchMot decoded mot %v", courierMot)

		switch courierMot.MotStatus {
		case domain.MOT_STATUS_NO_DETAILS_HELD_BY_DVLA:
			logger.Debugf("SearchMot no details held by dvla for registration number %v", in.GetRegistrationNumber())
			date, err := time.Parse(time.RFC3339, fmt.Sprintf("%s-01T11:45:26.371Z", courierMot.MonthOfFirstRegistration))
			if err != nil {
				logger.Errorf(fmt.Sprintf("SearchMot cannot parse month of first registration %v", courierMot.MonthOfFirstRegistration), err)
				return nil, proto.NoDetailsHeldByDvla.ErrorNoMsg()
			}
			logger.Debugf("SearchMot month of first registration is %v", date)
			if time.Now().Sub(date).Hours() > 24*365*10 {
				logger.Debugf("SearchMot month of first registration is more than one year ago")
				return nil, proto.NoDetailsHeldByDvla.ErrorNoMsg()
			} else {
				logger.Debugf("SearchMot month of first registration is less than one year ago")
			}
		case domain.MOT_STATUS_NO_RESULTS_RETURNED:
			logger.Debugf("SearchMot no results returned for registration number %v", in.GetRegistrationNumber())
			return nil, proto.NotFound.ErrorNoMsg()
		case domain.MOT_STATUS_NOT_VALID:
			logger.Debugf("SearchMot not valid registration number %v", in.GetRegistrationNumber())
			return nil, proto.InvalidRegistrationNumber.ErrorNoMsg()
		case domain.MOT_STATUS_VALID:
			logger.Debugf("SearchMot valid registration number %v", in.GetRegistrationNumber())
		}

		var markedForExport pb.Boolean
		if courierMot.MarkedForExport{
			markedForExport = pb.Boolean_BOOLEAN_TRUE
		} else {
			markedForExport = pb.Boolean_BOOLEAN_FALSE
		}

		return &pb.SearchMotResponse{
			Mot: &pb.Mot{
				RegistrationNumber:       courierMot.RegistrationNumber,
				Co2Emissions:             courierMot.Co2Emissions,
				EngineCapacity:           courierMot.EngineCapacity,
				EuroStatus:               courierMot.EuroStatus,
				MarkedForExport:          markedForExport,
				FuelType:                 courierMot.FuelType,
				MotStatus:                pb.MotStatus(courierMot.MotStatus),
				RevenueWeight:            courierMot.RevenueWeight,
				Colour:                   courierMot.Colour,
				Make:                     courierMot.Make,
				TypeApproval:             courierMot.TypeApproval,
				YearOfManufacture:        courierMot.YearOfManufacture,
				TaxDueDate:               courierMot.TaxDueDate,
				TaxStatus:                pb.TaxStatus(courierMot.TaxStatus),
				DateOfLastV5CIssued:      courierMot.DateOfLastV5CIssued,
				RealDrivingEmissions:     courierMot.RealDrivingEmissions,
				Wheelplan:                courierMot.Wheelplan,
				MonthOfFirstRegistration: courierMot.MonthOfFirstRegistration,
			},
		}, nil

	}

	logger.Debugf("SearchMot request to mot server is unknown")
	return nil, proto.Internal.ErrorMsg(fmt.Sprintf("response from server : %v", resp.StatusCode))
}

func (s *Service) GetMot(userId string) ([]*pb.Mot, error) {
	logger.Infof("GetMot userId = %v", userId)

	var motList []domain.CourierMot

	err := s.db.Model(&domain.CourierMot{}).Where("user_id = ?", userId).Scan(&motList).Error

	if err != nil {
		return nil, proto.Internal.Error(err)
	}

	if len(motList) == 0 {
		return nil, proto.NotFound.ErrorNoMsg()
	}

	var items = make([]*pb.Mot, len(motList))

	for i, item := range motList {
		var markedForExport pb.Boolean
		if item.MarkedForExport{
			markedForExport = pb.Boolean_BOOLEAN_TRUE
		} else {
			markedForExport = pb.Boolean_BOOLEAN_FALSE
		}
		items[i] = &pb.Mot{
			RegistrationNumber:       item.RegistrationNumber,
			Co2Emissions:             item.Co2Emissions,
			EngineCapacity:           item.EngineCapacity,
			EuroStatus:               item.EuroStatus,
			MarkedForExport:          markedForExport,
			FuelType:                 item.FuelType,
			MotStatus:                pb.MotStatus(item.MotStatus),
			RevenueWeight:            item.RevenueWeight,
			Colour:                   item.Colour,
			Make:                     item.Make,
			TypeApproval:             item.TypeApproval,
			YearOfManufacture:        item.YearOfManufacture,
			TaxDueDate:               item.TaxDueDate,
			TaxStatus:                pb.TaxStatus(item.TaxStatus),
			DateOfLastV5CIssued:      item.DateOfLastV5CIssued,
			RealDrivingEmissions:     item.RealDrivingEmissions,
			Wheelplan:                item.Wheelplan,
			MonthOfFirstRegistration: item.MonthOfFirstRegistration,
		}
	}

	logger.Infof("GetMot items = %v", items)

	return items, nil
}
