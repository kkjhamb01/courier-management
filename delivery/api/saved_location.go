package api

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/delivery/business"
	deliveryPb "github.com/kkjhamb01/courier-management/grpc/delivery/go"
	"github.com/kkjhamb01/courier-management/uaa/security"
)

func (s serverImpl) AddSavedLocation(ctx context.Context, req *deliveryPb.AddSavedLocationRequest) (*deliveryPb.AddSavedLocationResponse, error) {

	logger.Infof("AddSavedLocation request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate AddSavedLocationRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value("user").(security.User)

	var phone string
	if req.PhoneNumber != nil {
		phone = req.PhoneNumber.Value
	}

	var addressDetails string
	if req.AddressDetails != nil {
		addressDetails = req.AddressDetails.Value
	}

	var courierInstructions string
	if req.CourierInstructions != nil {
		courierInstructions = req.CourierInstructions.Value
	}

	id, err := business.AddSavedLocation(tokenUser.Id, req.Name, req.FullName, phone, addressDetails, req.Lat, req.Lon, courierInstructions)
	if err != nil {
		logger.Error("failed to add location", err)
		return nil, err
	}

	return &deliveryPb.AddSavedLocationResponse{
		Id: id,
	}, nil
}

func (s serverImpl) ListSavedLocations(ctx context.Context, req *empty.Empty) (*deliveryPb.ListSavedLocationsResponse, error) {

	logger.Infof("ListSavedLocations request = %+v", req)

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value("user").(security.User)

	locations, err := business.ListSavedLocations(tokenUser.Id)
	if err != nil {
		logger.Error("failed to list locations", err)
		return nil, err
	}

	return &deliveryPb.ListSavedLocationsResponse{
		SavedLocations: locations,
	}, nil
}

func (s serverImpl) UpdateSavedLocation(ctx context.Context, req *deliveryPb.UpdateSavedLocationRequest) (*empty.Empty, error) {

	logger.Infof("UpdateSavedLocation request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate UpdateSavedLocationRequest", err)
		return nil, err
	}

	err := business.UpdateSavedLocation(req.Id, req.Name, req.FullName, req.PhoneNumber, req.AddressDetails, req.Lat, req.Lon, req.CourierInstructions)
	if err != nil {
		logger.Error("failed to update location", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s serverImpl) RemoveSavedLocation(ctx context.Context, req *deliveryPb.RemoveSavedLocationRequest) (*empty.Empty, error) {

	logger.Infof("RemoveSavedLocation request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate RemoveSavedLocationRequest", err)
		return nil, err
	}

	err := business.RemoveSavedLocation(req.SavedLocationId)
	if err != nil {
		logger.Error("failed to remove location", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s serverImpl) RemoveAllSavedLocations(ctx context.Context, req *deliveryPb.RemoveAllSavedLocationsRequest) (*deliveryPb.RemoveAllSavedLocationsResponse, error) {

	logger.Infof("RemoveAllSavedLocations request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate RemoveAllSavedLocationsRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value("user").(security.User)

	locationIds, err := business.RemoveAllSavedLocations(tokenUser.Id)
	if err != nil {
		logger.Error("failed to remove all locations", err)
		return nil, err
	}

	return &deliveryPb.RemoveAllSavedLocationsResponse{
		SavedLocationIds: locationIds,
	}, nil
}
