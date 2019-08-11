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
