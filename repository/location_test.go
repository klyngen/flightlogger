package repository_test

import (
	"database/sql"
	"log"
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
	assert.Equal(t, "Øyer", location.AreaName)
	assert.Equal(t, "Norway", location.CountryName)

	// Assert that the location could be created
	assert.NilError(t, db.CreateLocation(&location))

	location.Name = "Kraabøl"

	assert.NilError(t, db.UpdateLocation(location.ID, &location))

	// Search should not be case sensitive
	locations, err := db.LocationSearchByName("RaA")

	// Assert that we did not have an error
	assert.NilError(t, err)

	// Assert that we got a result
	assert.Assert(t, len(locations) > 0)

	var getLocation common.Location
	assert.NilError(t, db.GetLocation(location.ID, &getLocation))

	// Lets assert the data to make sure that we still have not fucked anything up
	assert.Equal(t, "Kraabøl", location.Name)
	assert.Equal(t, getLocation.Lattitude, location.Lattitude)
	assert.Equal(t, getLocation.Longitude, location.Longitude)
	assert.Equal(t, getLocation.Elevation, location.Elevation)
	assert.Equal(t, getLocation.PostalCode, location.PostalCode)
	assert.Equal(t, getLocation.AreaName, location.AreaName)
	assert.Equal(t, getLocation.CountryName, location.CountryName)

	// Soft-delete the location
	assert.NilError(t, db.DeleteLocation(location.ID))

	// Try to get a deleted location
	var getDeletedLocation common.Location
	err = db.GetLocation(location.ID, &getDeletedLocation)

	assert.Assert(t, err == sql.ErrNoRows)
}

func TestStartSiteCycle(t *testing.T) {
	db := createConnection(t)

	location := common.Location{
		Name:        "Balbergkampen",
		Lattitude:   61.0,
		Longitude:   10.4662,
		Description: "Short PG-start with almost immediate lift after start",
		Elevation:   550,
		PostalCode:  "2624",
		AreaName:    "Lillehammer",
		CountryName: "Norway",
	}

	db.CreateLocation(&location)

	startsite := common.StartSite{
		Location:    location,
		Difficulty:  7,
		Description: "Reverse launch is a must",
		Name:        "",
	}

	assert.NilError(t, db.CreateStartSite(&startsite))
	log.Printf("ID of startsite %v", startsite.ID)

	// See if it has a sensible ID
	assert.Assert(t, startsite.ID > 0)

	startsite.Name = "Balbergkampen sørstart"

	sites, err := db.GetAllStartSites(10, 1)

	assert.NilError(t, err)
	assert.Assert(t, len(sites) > 0)

	assert.NilError(t, db.UpdateStartSite(startsite.ID, &startsite))

	var site common.StartSite
	err = db.GetStartSite(startsite.ID, &site)
	log.Println(site.ID)
	assert.NilError(t, err)

	// Now see that the data is correct
	assert.Equal(t, startsite.Name, site.Name)
	assert.Equal(t, startsite.Difficulty, site.Difficulty)
	assert.Equal(t, startsite.Description, site.Description)

	_, err = db.GetStartSiteWaypoints(startsite.ID)

	assert.NilError(t, err)

	// The cherry on top is to soft delete this stuff
	assert.NilError(t, db.DeleteStartSite(startsite.ID))
	assert.NilError(t, db.DeleteLocation(location.ID))

}

func TestWaypointCycle(t *testing.T) {
	db := createConnection(t)

	location := common.Location{
		Name:        "Bangsberget",
		Lattitude:   60.9,
		Longitude:   10.4662,
		Description: "Short PG-start with almost immediate lift after start",
		Elevation:   600,
		PostalCode:  "2600",
		AreaName:    "Brummundal",
		CountryName: "Norway",
	}

	db.CreateLocation(&location)

	waypoint := common.Waypoint{
		Location:    location,
		Difficulty:  7,
		Description: "Good start-site heading east",
		Name:        "",
	}

	assert.NilError(t, db.CreateWayPoint(&waypoint))
	log.Printf("ID of startsite %v", waypoint.ID)

	waypoint.Name = "Bangsjordet"
	// See if it has a sensible ID
	assert.Assert(t, waypoint.ID > 0)

	assert.NilError(t, db.UpdateWayPoint(waypoint.ID, &waypoint))

	sites, err := db.GetAllWaypoints(10, 0)
	assert.NilError(t, err)

	assert.Assert(t, len(sites) > 0)

	var site common.Waypoint
	err = db.GetWaypoint(waypoint.ID, &site)
	log.Println(site.ID)
	assert.NilError(t, err)

	// Now see that the data is correct
	assert.Equal(t, waypoint.Name, site.Name)
	assert.Equal(t, waypoint.Difficulty, site.Difficulty)
	assert.Equal(t, waypoint.Description, site.Description)

	// The cherry on top is to soft delete this stuff
	assert.NilError(t, db.DeleteWaypoint(waypoint.ID))
	assert.NilError(t, db.DeleteLocation(location.ID))
}
