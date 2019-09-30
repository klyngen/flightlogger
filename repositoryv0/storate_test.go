package storage

import (
	"log"
	"os"
	"testing"

	"gotest.tools/assert"

	"github.com/klyngen/flightlogger/common"
)

func SetupDbTest() *OrmDatabase {
	var db = &OrmDatabase{}

	err := db.CreateConnection("root", "", "flightlog", "3306", "localhost")

	if err != nil {
		log.Println(err)
	}

	return db
}

func TestMigration(t *testing.T) {
	db := SetupDbTest()
	db.db.LogMode(true)
	db.db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	err := db.MigrateDatabase()

	if err != nil {
		t.Fatalf("Migration failed %v", err)
	}
}

func TestUserLifeCycle(t *testing.T) {
	db := SetupDbTest()

	user := common.User{
		Username:     "klyngen",
		LastName:     "klingenberg",
		FirstName:    "Martin",
		Email:        "martin@klingenberg.as",
		PasswordHash: []byte("somehash"),
		PasswordSalt: []byte("something salty"),
	}

	// Create a user
	newUser, err := db.CreateUser(user)

	if err != nil {
		t.Fatalf("Could not create database user with the following error %v", err)
	}

	// Assert that the parameters are still the same after creation
	assert.Equal(t, user.Username, newUser.Username)
	assert.Equal(t, user.LastName, newUser.LastName)
	assert.Equal(t, user.FirstName, newUser.FirstName)
	assert.Equal(t, user.Email, newUser.Email)

	// Fix the typo in the name
	user.LastName = "Klingenberg"

	var updatedUser common.User

	updatedUser, err = db.UpdateUser(newUser.ID, user)

	if err != nil {
		t.Fatalf("Could not update database user with the following error %v", err)
	}

	// Names should no longer be the same
	assert.Assert(t, newUser.LastName != updatedUser.LastName)

	var storedUser common.User

	storedUser, err = db.GetUser(newUser.ID)

	if err != nil {
		t.Fatalf("Could not get database user with the following error %v", err)
	}

	assert.Equal(t, newUser.ID, storedUser.ID)

	// Complete by deleting the user
	err = db.DeleteUser(newUser.ID)

	if err != nil {
		t.Fatalf("Could not delete database user with the following error %v", err)
	}

}

func TestLocationCycle(t *testing.T) {
	db := SetupDbTest()
	location := common.Location{
		Name:        "Gjelle",
		Lattitude:   61.02,
		Longitude:   61.5,
		Description: "Small town in the west of norway. Voss is known for its love for extreme sports",
		PostalCode:  "1",
		AreaName:    "Oslo",
		CountryPart: "Oslo",
	}

	location2 := common.Location{
		Name:        "Balbergtoppen",
		Lattitude:   61.02,
		Longitude:   61.5,
		Description: "Not really in Oslo",
		PostalCode:  "1",
		AreaName:    "Oslo",
		CountryPart: "Oslo",
	}

	newLocation, err := db.CreateLocation(location)

	if err != nil {
		t.Failed()
	}

	// Test if we can store a exact similar location without
	_, err = db.CreateLocation(location2)

	if err != nil {
		t.Fatalf("Failed to store with existing countryPart: %v", err)
	}

	assert.Equal(t, newLocation.Description, location.Description)
	assert.Equal(t, newLocation.Name, location.Name)
	assert.Equal(t, newLocation.Lattitude, location.Lattitude)
	assert.Equal(t, newLocation.Longitude, location.Longitude)

	newLocation.Longitude = 70

	var updated common.Location
	updated, err = db.UpdateLocation(newLocation.ID, newLocation)

	if err != nil {
		t.Fatalf("Unable to update location: %v", err)
	}

	// Assert that we have actually updated the location
	assert.Assert(t, updated.Longitude != location.Longitude)

	searchResult, err := db.LocationSearchByName("gj")

	if err != nil {
		t.Fatalf("Unable to search for locations: %v", err)
	}

	// We should have one search-result from the database
	assert.Assert(t, 0 < len(searchResult))

	// Try deletion
	err = db.DeleteLocation(newLocation.ID)

	if err != nil {
		t.Fatalf("Cannot delete a location: %v", err)
	}

}

func TestStartSiteWaypointCycle(t *testing.T) {
	db := SetupDbTest()
	location := common.Location{
		Name:        "Gjelle",
		Lattitude:   61.02,
		Longitude:   61.5,
		Description: "Small town in the west of norway. Voss is known for its love for extreme sports",
		PostalCode:  "1",
		AreaName:    "Oslo",
		CountryPart: "Oslo",
	}

	loc, err := db.CreateLocation(location)

	wps := createWayPoints(t, db)

	startSite := common.StartSite{
		Difficulty: 5,
		Location:   loc,
		Waypoints: []common.Waypoint{
			common.Waypoint{ID: wps[0]},
		},
	}

	newStartSite, err := db.CreateStartSite(startSite)

	if err != nil {
		t.Fatalf("Unable to create startSite %v", err)
	}
	// Verify mapping
	assert.Assert(t, newStartSite.ID != 0)
	assert.Equal(t, newStartSite.Difficulty, startSite.Difficulty)
	assert.Equal(t, newStartSite.Description, startSite.Description)
	// Do we still have one waypoint?
	assert.Equal(t, 1, len(newStartSite.Waypoints))

	// We should be able to replace one waypoint with another
	newStartSite.Waypoints = []common.Waypoint{common.Waypoint{ID: wps[1]}}

	updatedSite, err := db.UpdateStartSite(newStartSite.ID, newStartSite)

	// See if the correct waypoint is in the correct place
	assert.Assert(t, updatedSite.Waypoints[0].ID != wps[0])
	// See that we only have one waypoint
	assert.Equal(t, 1, len(updatedSite.Waypoints))

	updateWaypoints(wps, t, db)

	// Cleanup
	err = deleteWaypoints([]uint{wps[0], wps[1]}, t, db)

	if err != nil {
		t.Fatalf("Could not clean up waypoints %v", err)
	}

}

func createWayPoints(t *testing.T, database *OrmDatabase) []uint {
	location := common.Location{
		Name:        "Gjelle landing",
		Lattitude:   61.02,
		Longitude:   61.5,
		Description: "Landing ved Gjelle. Følg vannet",
		PostalCode:  "1",
		AreaName:    "Oslo",
		CountryPart: "Oslo",
	}

	location2 := common.Location{
		Name:        "Gjelle landing2",
		Lattitude:   61.02,
		Longitude:   61.5,
		Description: "Landing på feil side av høyspentlinjen",
		PostalCode:  "1",
		AreaName:    "Oslo",
		CountryPart: "Oslo",
	}

	// Is not stored in the database
	location3 := common.Location{
		Name:        "Not valid",
		Lattitude:   0.0,
		Longitude:   0.0,
		Description: "Landing på feil side av høyspentlinjen",
		PostalCode:  "1",
		AreaName:    "Oslo",
		CountryPart: "Oslo",
	}

	dbLocation, err := database.CreateLocation(location)
	dbLocation2, err := database.CreateLocation(location2)

	if err != nil {
		t.Fatalf("Unable to make locations for waypoints... %v", err)
	}

	wp1 := common.Waypoint{
		Difficulty: 5,
		Location:   dbLocation,
	}

	wp2 := common.Waypoint{
		Difficulty: 5,
		Location:   dbLocation2,
	}

	wp3 := common.Waypoint{
		Difficulty: 1,
		Location:   location3,
	}

	waypoint1, err := database.CreateWayPoint(wp1)
	waypoint2, err := database.CreateWayPoint(wp2)

	if err != nil {
		t.Fatalf("Unable to store waypoints %v", err)
	}

	_, err = database.CreateWayPoint(wp3)

	// Then the waypoint was created without a valid location
	if err == nil {
		t.Fatalf("Waypoint created without location %v", err)
	}

	return []uint{waypoint1.ID, waypoint2.ID}
}

func updateWaypoints(waypoints []uint, t *testing.T, d *OrmDatabase) error {
	for _, p := range waypoints {
		var wp DbWaypoint
		d.db.First(&wp, p)

		mappedPoint := demapWaypoint(wp)
		// It is suddenly really easy to land here
		wp.Difficulty = 1

		_, err := d.UpdateWayPoint(p, mappedPoint)

		if err != nil {
			t.Fatalf("Unable to delete waypoints %v", err)
		}
	}
	return nil
}

func deleteWaypoints(waypoints []uint, t *testing.T, d *OrmDatabase) error {
	for _, p := range waypoints {
		err := d.DeleteWayPoint(p)

		if err != nil {
			t.Fatalf("Unable to delete waypoints %v", err)
		}
	}
	return nil
}

func TestFlightCrud(t *testing.T) {
	db := SetupDbTest()
	user := common.User{
		Username:     "klyngen",
		LastName:     "klingenberg",
		FirstName:    "Martin",
		Email:        "martin@klingenberg.as",
		PasswordHash: []byte("somehash"),
		PasswordSalt: []byte("something salty"),
	}

	storedUser, err := db.CreateUser(user)

	flight := common.Flight{}

	db.DeleteUser(storedUser.ID)
}
