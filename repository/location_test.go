package repository_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/klyngen/flightlogger/common"
)

func TestLocationCreation(t *testing.T) {
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

	// Assert that the location could be created
	assert.NilError(t, db.CreateLocation(&location))
}
