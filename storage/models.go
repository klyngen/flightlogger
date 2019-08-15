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
	Username  string
	FirstName string
	LastName  string
	Email     string        `gorm:"type:varchar(100);unique_index"`
	Clubs     []DbClub      `gorm:"many2many:user_clubs;"`
	Scopes    []DbUserScope `gorm:"many2many:user_scopes;"`
	Groups    []DbUserGroup `gorm:"many2many:user_group;"`
	Wings     []DbWing      `gorm:"many2many:user_wing;"`
}

// DbCredentials - the login credentials for a user
type DbCredentials struct {
	gorm.Model
	UserID       uint
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
	Coordinates        DbCoordinates `gorm:"foreignkey:CoordinatesReferer"`
	CountryPart        DbCountryPart `gorm:"foreignkey:CountrypartReferer"`
	CountrypartReferer uint
	CoordinatesReferer uint
	Name               string
	Description        string
	Elevation          int
}

// DbCountryPart - describes the part of the country example Lillehammer, Oppland 2620
// the primary key is described by a unique combination of the three parameters
type DbCountryPart struct {
	gorm.Model
	AreaName    string `gorm:"primary_key"`
	PostalCode  string `gorm:"primary_key"`
	CountryPart string `gorm:"primary_key"`
}

// DbStartSite - describes a start sight for flight
type DbStartSite struct {
	gorm.Model
	Waypoints       []DbWaypoint `gorm:"many2many:startsite_waypoints;"`
	Location        DbLocation
	LocationReferer uint `gorm:"foreignkey:LocationReferer"`
	Difficulty      int
}

// DbWaypoint - describes a start sight for flight
type DbWaypoint struct {
	gorm.Model
	Difficulty      int
	LocationReferer uint `gorm:"foreignkey:LocationReferer"`
	Location        DbLocation
}

// DbFlight - describes a flight
type DbFlight struct {
	gorm.Model
	Startsite   DbStartSite `gorm:"association_foreignkey:Refer"`
	User        DbUser      `gorm:"association_foreignkey:Refer"`
	Waypoint    DbWaypoint  `gorm:"association_foreignkey:Refer"`
	Duration    int
	Notes       string
	Distance    int
	MaxHight    int
	HangTime    int
	Wing        DbWing `gorm:"many2many:flight_wing;"`
	Incidents   []DbIncident
	Photos      []DbFileReference
	FlightLog   DbFileReference `gorm:"association_foreignkey:Refer"`
	FlightType  DbFlightType    `gorm:"association_foreignkey:Refer"`
	TakeOffType DbTakeoffType   `gorm:"association_foreignkey:Refer"`
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

// ################# MODEL HELPER FUNCTIONS #########################################

// IsEmpty returns true if no parameter has a value
func (cp *DbCountryPart) isEmpty() bool {
	if len(cp.AreaName) == 0 && len(cp.CountryPart) == 0 && len(cp.PostalCode) == 0 {
		return true
	}
	return false
}
