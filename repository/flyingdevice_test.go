package repository_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/klyngen/flightlogger/common"
)

func TestFlyingDeviceCreation(t *testing.T) {
	database := createConnection(t)

	device := common.FlyingDevice{
		Model:      "Alpha 6",
		Make:       "Advane", // Bad spelling is not a mistake in this situation
		Details:    make([]common.FlyingDeviceDetails, 0),
		DeviceType: 3,
	}

	device.Details = append(device.Details, common.FlyingDeviceDetails{
		Description: "EN/A",
		DetailName:  "EN-Rating",
	})

	assert.NilError(t, database.CreateWing(&device), "Could not create flyingDevice")
	assert.Assert(t, device.ID > 0) // Check that we have an ID

	device.Make = "Advance"

	assert.NilError(t, database.UpdateWing(device.ID, &device), "Could not update the flyingDevice")

	// Try getting the device in various ways
	var newDevice common.FlyingDevice
	assert.NilError(t, database.GetWing(device.ID, &newDevice), "Could not fetch an existing wing")

	// Make sure the raw data is good
	assert.Equal(t, device.ID, newDevice.ID)
	assert.Equal(t, device.Model, newDevice.Model)
	assert.Equal(t, device.Make, newDevice.Make)
	assert.Assert(t, len(newDevice.Details) > 0)

	rows, err := database.GetAllWings(10, 1)

	assert.NilError(t, err, "Could not fetch multiple devices")

	assert.Assert(t, len(rows) > 0)

	assert.NilError(t, database.DeleteWing(device.ID), "Could not delete the device")
}
