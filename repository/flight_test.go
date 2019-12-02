package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/klyngen/flightlogger/common"
)

func TestFlightCreationCycle(t *testing.T) {
	db := createConnection(t)

	// Create a user
	user := common.User{
		FirstName:    "Martin",
		LastName:     "Klingenberg",
		Email:        "m@m.no",
		PasswordHash: []byte("hash"),
	}

	db.CreateUser(&user)
	defer db.DeleteUser(user.ID)

	location := common.Location{
		Name:        "Bangsberget",
		Lattitude:   10.0,
		Longitude:   60.0,
		CountryName: "Norway",
		PostalCode:  "1000",
		AreaName:    "Voss",
	}

	db.CreateLocation(&location)
	defer db.DeleteLocation(location.ID)

	startsite := common.StartSite{
		Name:       "Bangsberg Øst-start",
		Location:   location,
		Difficulty: 6,
	}

	db.CreateStartSite(&startsite)
	defer db.DeleteStartSite(startsite.ID)

	wing := common.FlyingDevice{
		Model:      "Alpha 6",
		Make:       "Advance",
		DeviceType: 1,
	}

	db.CreateWing(&wing)
	defer db.DeleteWing(wing.ID)

	// Now do the actual testing....
	flight := common.Flight{
		User:      user,
		Startsite: startsite,
		Duration:  20,
		Notes:     "Should be a long explanation",
		Wing:      wing,
	}

	assert.NoError(t, db.CreateFlight(&flight))

	flight.Duration = 40
	flight.HangTime = 20

	assert.NoError(t, db.UpdateFlight(flight.ID, &flight))
	// CLEAN UP THE FLIGHT
	assert.NoError(t, db.DeleteFlight(flight.ID, true))
	// SHOULD BE ABLE TO BOTH SOFT AND HARD-DELETE
	assert.NoError(t, db.DeleteFlight(flight.ID, false))

}
