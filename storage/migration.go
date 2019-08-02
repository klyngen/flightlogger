package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/klyngen/flightlogger/common"
)

/*
	THIS FILE IS RESPONSIBLE FOR THE MIGRATION PROCESS OF THE DATABASE
	REMEMBER HOW THE CREATION PROCESS HAS A SPECIFIC ORDER
*/

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	// Migrate location first
	db.AutoMigrate(&common.DbFileReference{})
	db.AutoMigrate(&common.DbCoordinates{})
	db.AutoMigrate(&common.DbLocation{})

	// Create club entity before user and flights
	db.AutoMigrate(&common.DbClub{})

	// Waypoint and start are dependent on location
	db.AutoMigrate(&common.DbWaypoint{})
	db.AutoMigrate(&common.DbStartSite{})

	// Wing related data
	db.AutoMigrate(&common.DbWingScoreDetails{})
	db.AutoMigrate(&common.DbWing{})

	// Flight related entities
	db.AutoMigrate(&common.DbFlightType{})
	db.AutoMigrate(&common.DbTakeoffType{})
	db.AutoMigrate(&common.DbIncident{})
	db.AutoMigrate(&common.DbFlight{})

	// Set up the user related entities
	db.AutoMigrate(&common.DbCredentials{})
	db.AutoMigrate(&common.DbUserScope{})
	db.AutoMigrate(&common.DbUserGroup{})
	db.AutoMigrate(&common.DbUser{})
	return db
}
