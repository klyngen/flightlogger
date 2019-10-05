package repository_test

import (
	"database/sql"
	"testing"

	"gotest.tools/assert"

	"github.com/klyngen/flightlogger/common"
)

func TestLocationCreationCycle(t *testing.T) {
	location := common.Location{
		Name:        "Kråbøl",
		Lattitude:   61.0,
		Longitude:   10.4662,
		Description: "Historically been a HG start",
		Elevation:   652,
		PostalCode:  "2636",
		CountryPart: "Oppland",
		AreaName:    "Øyer",
		CountryName: "Norway",
	}

	db := createConnection(t)

	// Lets assert the data to make sure we have not fucked anything up
	assert.Equal(t, "Kråbøl", location.Name)
	assert.Equal(t, 61.0, location.Lattitude)
	assert.Equal(t, 10.4662, location.Longitude)
	assert.Equal(t, 652, location.Elevation)
	assert.Equal(t, "2636", location.PostalCode)
	assert.Equal(t, "Oppland", location.CountryPart)
	assert.Equal(t, "Norway", location.CountryName)

	// Assert that the location could be created
	assert.NilError(t, db.CreateLocation(&location))

	location.Name = "Kraabøl"

	assert.NilError(t, db.UpdateLocation(location.ID, &location))

	var getLocation common.Location
	assert.NilError(t, db.GetLocation(location.ID, &getLocation))

	// Lets assert the data to make sure that we still have not fucked anything up
	assert.Equal(t, "Kraabøl", location.Name)
	assert.Equal(t, getLocation.Lattitude, location.Lattitude)
	assert.Equal(t, getLocation.Longitude, location.Longitude)
	assert.Equal(t, getLocation.Elevation, location.Elevation)
	assert.Equal(t, getLocation.PostalCode, location.PostalCode)
	assert.Equal(t, getLocation.CountryPart, location.CountryPart)
	assert.Equal(t, getLocation.CountryName, location.CountryName)

	// Soft-delete the location
	assert.NilError(t, db.DeleteLocation(location.ID))

	// Try to get a deleted location
	var getDeletedLocation common.Location
	assert.Assert(t, db.GetLocation(location.ID, &getDeletedLocation) == sql.ErrNoRows)
}
