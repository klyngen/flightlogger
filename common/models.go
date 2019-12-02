package common

import (
	"time"
)

/*
	THIS FILE CONTAINS A SIMPLIFIED SET OF MODELS EXPOSED TO THE API ENDPOINTS
	THEESE MODELS SHOULD BE TOTALLY INDEPENDENT FROM THE DATA LAYER
	THE USAGE SHOULD BE PRIMARILY USED BY THE PRESENTATION AND BUSINESS LAYER
*/

// User - defines the basic user model exposed on the API's
type User struct {
	ID            string
	Username      string
	FirstName     string
	LastName      string
	Email         string
	Active        bool `json:"-"`
	Clubs         []Club
	Scopes        []UserScope
	Groups        []UserGroup
	Wings         []FlyingDevice
	TimeUpdated   time.Time
	TimeGenerated time.Time
	PasswordHash  []byte `json:"-"` // Salt and hash should not be a part of serialized JSON
}

// UserScope - the possible scopes of a user
type UserScope struct {
	ID        uint
	Key       string
	Name      string
	TimeStamp time.Time
}

// UserGroup - defines a set of scopes that can be applied to a user
type UserGroup struct {
	ID     uint
	Key    string
	Name   string
	Scopes []UserScope
}

// Club - describes a paragliding club
type Club struct {
	ID          uint
	Location    Location
	Description string
	Name        string
}

// Location - describes a place
type Location struct {
	ID          uint
	Name        string
	Elevation   int
	Description string
	Longitude   float64
	Lattitude   float64
	CountryName string
	AreaName    string
	PostalCode  string
}

// StartSite - describes a start sight for flight
type StartSite struct {
	Location
	ID          uint
	Name        string
	Waypoints   []Waypoint
	Difficulty  int
	Description string
}

// Waypoint - describes a start sight for flight
type Waypoint struct {
	Location
	ID          uint
	Difficulty  int
	Name        string
	Description string
}

// Flight - describes a flight
type Flight struct {
	ID          string
	User        User
	Startsite   StartSite
	Waypoint    Waypoint
	Duration    int
	Notes       string
	Distance    int
	MaxHight    int
	HangTime    int
	Wing        FlyingDevice
	Incidents   []Incident
	Photos      []FileReference
	FlightLog   FileReference
	FlightType  FlightType
	TakeOffType TakeoffType
	Created     time.Time
}

// Incident - describes an incident
type Incident struct {
	ID                   uint
	Level                int
	Description          string
	Public               bool
	NotifiedPolice       bool
	NotifiedAmbulance    bool
	AirAmbulance         bool
	HeadOfSecurityCalled bool
	LatestFlightID       uint
	Weatherconfitions    string
}

// Wing - describes a wing
type FlyingDevice struct {
	ID         uint
	Model      string
	Make       string
	Images     []FileReference
	DeviceType FlyingDeviceType
	Details    []FlyingDeviceDetails
}

// WingDetails - details of the wing test
type FlyingDeviceDetails struct {
	ID          uint
	Description string
	DetailName  string
}

// FileReference - describes whre a file is stored
type FileReference struct {
	ID           uint
	MimeType     string
	FileName     string
	FileLocation string
}

// TakeoffType - describes a type of takeoff
type TakeoffType struct {
	Name  string
	Value uint
}

// FlightType - is this a Paragliding, Speedrider, Baloon etc
type FlightType struct {
	Name  string
	Value uint
}
