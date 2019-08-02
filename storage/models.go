package storage

import (
	"time"

	"github.com/jinzhu/gorm"
)

/*
 * THIS FILE CONTAINS DATABASE STRUCTS ALL STRUCTS SHOULD BE PREFIX WITH DB
 */

// DbUser - the ground pillar of the userconstruct
type DbUser struct {
	gorm.Model
	Username    string
	FirstName   string
	LastName    string
	Email       string   `gorm:"type:varchar(100);unique_index"`
	Clubs       []DbClub `gorm:"many2many:user_scopes;"`
	Credentials DbCredentials
	Scopes      []DbUserScope `gorm:"many2many:user_scopes;"`
	Groups      []DbUserGroup `gorm:"many2many:user_group;"`
	Wings       []DbWing      `gorm:"many2many:user_wing;"`
}

// DbCredentials - the login credentials for a user
type DbCredentials struct {
	gorm.Model
	PasswordHash []byte
	PasswordSalt []byte
}

// DbUserScope - the possible scopes of a user
type DbUserScope struct {
	gorm.Model
	Key       string
	Name      string
	TimeStamp time.Time
}

// DbUserGroup - defines a set of scopes that can be applied to a user
type DbUserGroup struct {
	gorm.Model
	Key    string
	Name   string
	Scopes []DbUserScope
}

// DbClub - describes a paragliding club
type DbClub struct {
	gorm.Model
	Location    DbLocation
	Description string
	Name        string
}

// DbCoordinates - describes a set of coordinates
type DbCoordinates struct {
	gorm.Model
	Longitude float64
	Lattitude float64
}

// DbLocation - describes a place
type DbLocation struct {
	gorm.Model
	Coordinates DbCoordinates
	Name        string
	Description string
	Elevation   int
}

// DbStartSite - describes a start sight for flight
type DbStartSite struct {
	gorm.Model
	Description string
	Waypoints   []DbWaypoint
	Location    DbLocation
	Difficulty  int
}

// DbWaypoint - describes a start sight for flight
type DbWaypoint struct {
	gorm.Model
	Difficulty int
	Location   DbLocation
}

// DbFlight - describes a flight
type DbFlight struct {
	gorm.Model
	Startsite   DbStartSite
	User        DbUser
	Waypoint    DbWaypoint
	Duration    int
	Notes       string
	Distance    int
	MaxHight    int
	HangTime    int
	Wing        DbWing `gorm:"many2many:flight_wing;"`
	Incidents   []DbIncident
	Photos      []DbFileReference
	FlightLog   DbFileReference
	FlightType  DbFlightType
	TakeOffType DbTakeoffType
}

// DbIncident - describes an incident
type DbIncident struct {
	gorm.Model
	Level                int
	Description          string
	Public               bool
	NotifiedPolice       bool
	NotifiedAmbulance    bool
	AirAmbulance         bool
	HeadOfSecurityCalled bool
	LatestFlight         DbFlight
	Weatherconfitions    string
	// TODO: extend to better reflect the original flightlog
}

// DbWing - describes a wing
type DbWing struct {
	gorm.Model
	Name    string
	Details DbWingScoreDetails
	Images  []DbFileReference
}

// DbWingScoreDetails - describes which rating the wing has EN-A, EN-B
type DbWingScoreDetails struct {
	gorm.Model
	Description string
	DhvScore    string
	EnaScore    string
}

// DbFileReference - describes whre a file is stored
type DbFileReference struct {
	gorm.Model
	MimeType     string
	FileName     string
	FileLocation string
}

// THE LAST TWO CAN HAVE SOFT DELETE BECAUSE THEY ARE CONTSTANT AND
// WONT VIOLATE GDPR

// DbTakeoffType - describes a type of takeoff
type DbTakeoffType struct {
	gorm.Model
	Name string
}

// DbFlightType - is this a Paragliding, Speedrider, Baloon etc
type DbFlightType struct {
	gorm.Model
	Name string
}
