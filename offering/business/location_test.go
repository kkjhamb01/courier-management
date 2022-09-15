package business

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetCourierLocationEmptyLocation(t *testing.T) {
	err := SetCourierLocation(context.Background(), offeringPb.Location{}, offeringPb.CourierType_TRUCK, "18da79e8-6b60-4d35-a3c9-be28ff9e731d")
	assert.NotNil(t, err, "an error is expected when the location is empty")
}

func TestSetCourierLocationInvalidLocationLat(t *testing.T) {
	err := SetCourierLocation(context.Background(), offeringPb.Location{
		Lat: 92, //max: 90
		Lon: 120,
	},
		offeringPb.CourierType_TRUCK, "18da79e8-6b60-4d35-a3c9-be28ff9e731d")
	assert.NotNil(t, err, "an error is expected when the location lat is not valid")
}

func TestSetCourierLocationInvalidLocationLon(t *testing.T) {
	err := SetCourierLocation(context.Background(), offeringPb.Location{
		Lat: 82,
		Lon: -190, //min: -180
	},
		offeringPb.CourierType_TRUCK, "18da79e8-6b60-4d35-a3c9-be28ff9e731d")
	assert.NotNil(t, err, "an error is expected when the location lon is not valid")
}
