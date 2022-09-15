package business

import (
	"errors"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/delivery/db"
	"github.com/kkjhamb01/courier-management/delivery/model"
	deliveryPb "github.com/kkjhamb01/courier-management/grpc/delivery/go"
	uaa "github.com/kkjhamb01/courier-management/uaa/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func AddSavedLocation(customerId, name, fullName, phoneNumber string, addressDetails string, lat, lon float64, courierInstructions string) (string, error) {

	logger.Infof("AddSavedLocation customerId = %v, name = %v, fullName = %v, phoneNumber = %v, addressDetails = %v, lat = %v, lon = %v, courierInstructions = %v", customerId, name, fullName, phoneNumber, addressDetails, lat, lon, courierInstructions)

	savedLocation := model.SavedLocation{
		CustomerId:          customerId,
		Name:                name,
		FullName:            fullName,
		PhoneNumber:         &phoneNumber,
		AddressDetails:      &addressDetails,
		Lat:                 lat,
		Lon:                 lon,
		CourierInstructions: &courierInstructions,
	}
	result := db.MariaDbClient().Create(&savedLocation)
	if result.Error != nil {
		logger.Error("failed to create savedLocation", result.Error)
		return "", uaa.Internal.Error(result.Error)
	}

	if result.RowsAffected == 0 {
		err := errors.New("no row is affected")
		logger.Error("failed to create savedLocation", err)
		return "", uaa.NotFound.Error(err)
	}

	return savedLocation.ID, nil
}

func ListSavedLocations(customerId string) ([]*deliveryPb.SavedLocation, error) {
	var addresses []*model.SavedLocation

	result := db.MariaDbClient().Where(map[string]interface{}{"customer_id": customerId}).Find(&addresses)
	if result.Error != nil {
		logger.Error("failed to list addresses", result.Error)
		return nil, uaa.Internal.Error(result.Error)
	}

	var protoAddresses = make([]*deliveryPb.SavedLocation, len(addresses))
	for i, address := range addresses {
		protoAddresses[i] = address.ToProtoP()
	}

	return protoAddresses, nil
}

func UpdateSavedLocation(addressId string, name, fullName, phoneNumber, addressDetails *wrapperspb.StringValue, lat, lon *wrapperspb.DoubleValue, courierInstructions *wrapperspb.StringValue) error {
	fields := make(map[string]interface{}, 7)

	if name != nil {
		fields["name"] = name.Value
	}
	if fullName != nil {
		fields["full_name"] = fullName.Value
	}
	if phoneNumber != nil {
		fields["phone_number"] = phoneNumber.Value
	}
	if addressDetails != nil {
		fields["address_details"] = addressDetails.Value
	}
	if lat != nil {
		fields["lat"] = lat.Value
	}
	if lon != nil {
		fields["lon"] = lon.Value
	}
	if courierInstructions != nil {
		fields["courier_instructions"] = courierInstructions.Value
	}

	address := &model.SavedLocation{}
	address.ID = addressId
	result := db.MariaDbClient().Model(address).Updates(fields)
	if result.Error != nil {
		logger.Error("failed to update address", result.Error)
		return uaa.Internal.Error(result.Error)
	}

	return nil
}

func RemoveSavedLocation(addressId string) error {
	result := db.MariaDbClient().Where(map[string]interface{}{"Id": addressId}).Delete(&model.SavedLocation{})
	if result.Error != nil {
		logger.Error("failed to delete address", result.Error)
		return uaa.Internal.Error(result.Error)
	}

	if result.RowsAffected == 0 {
		err := errors.New("now row affected")
		logger.Error("failed to delete address", err)
		return uaa.NotFound.Error(err)
	}

	return nil
}

func RemoveAllSavedLocations(customerId string) ([]string, error) {
	var addresses []*model.SavedLocation

	queryResult := db.MariaDbClient().
		Where(map[string]interface{}{"customer_id": customerId}).
		Select("ID").
		Find(&addresses)
	if queryResult.Error != nil {
		logger.Error("failed to list address ids", queryResult.Error)
		return nil, uaa.Internal.Error(queryResult.Error)
	}

	result := db.MariaDbClient().Delete(addresses)
	if result.Error != nil {
		logger.Error("failed to delete addresses", result.Error)
		return nil, uaa.Internal.Error(result.Error)
	}

	if result.RowsAffected == 0 {
		err := errors.New("now row affected")
		logger.Error("failed to delete addresses", err)
		return nil, uaa.NotFound.Error(err)
	}

	addressesIds := make([]string, len(addresses))
	for i, address := range addresses {
		addressesIds[i] = address.ID
	}

	return addressesIds, nil
}
