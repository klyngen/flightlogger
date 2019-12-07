package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/klyngen/flightlogger/common"
)

func resolveUser(db common.FlightLogDatabase) common.User {
	var user common.User

	email := "m@m.no"

	db.GetUserByEmail(email, &user)

	if len(user.Email) > 0 {
		return user
	}

	user = common.User{
		FirstName:    "Martin",
		LastName:     "Klingenberg",
		Email:        email,
		PasswordHash: []byte("hash"),
		Active:       true,
	}

	db.CreateUser(&user)

	return user
}

func TestFlightCreationCycle(t *testing.T) {
	db := createConnection(t)

	// Create a user
	user := resolveUser(db)

	location := common.Location{
		Name:        "Hangurtoppen",
		Lattitude:   10.0,
		Longitude:   60.0,
		CountryName: "Norway",
		PostalCode:  "1000",
		AreaName:    "Voss",
	}

	db.CreateLocation(&location)

	startsite := common.StartSite{
		Name:       "Hangurtoppen Øst-start",
		Location:   location,
		Difficulty: 6,
	}

	db.CreateStartSite(&startsite)

	wing := common.FlyingDevice{
		Model:      "Alpha 6",
		Make:       "Advance",
		DeviceType: 1,
	}

	db.CreateWing(&wing)

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

	var newFlight common.Flight

	assert.NoError(t, db.GetFlight(flight.ID, &newFlight))

	assert.Equal(t, flight.User.ID, newFlight.User.ID)
	assert.Equal(t, flight.Startsite.ID, newFlight.Startsite.ID)
	assert.Equal(t, flight.Wing.ID, newFlight.Wing.ID)
	assert.Equal(t, flight.Distance, newFlight.Distance)
	assert.Equal(t, flight.Duration, newFlight.Duration)
	assert.Equal(t, flight.MaxHight, newFlight.MaxHight)

	// CLEAN UP THE FLIGHT
	assert.NoError(t, db.DeleteFlight(flight.ID, true))
	// SHOULD BE ABLE TO BOTH SOFT AND HARD-DELETE
	assert.NoError(t, db.DeleteFlight(flight.ID, false))

	db.DeleteUser(user.ID)
	db.DeleteWing(wing.ID)
	db.DeleteStartSite(startsite.ID)
	db.DeleteLocation(location.ID)
}
