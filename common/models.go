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
	ID            int
	Username      string
	FirstName     string
	LastName      string
	Email         string
	Clubs         []Club
	Scopes        []UserScope
	Group         []UserGroup
	Wings         []Wing
	PasswordHash  []byte
	PasswordSalt  []byte
	TimeUpdated   time.Time
	TimeGenerated time.Time
}

// UserScope - the possible scopes of a user
type UserScope struct {
	ID        int
	Key       string
	Name      string
	TimeStamp time.Time
}

// UserGroup - defines a set of scopes that can be applied to a user
type UserGroup struct {
	ID    int
	Key   string
	Name  string
	Scope []UserScope
}

// Club - describes a paragliding club
type Club struct {
	ID          int
	Location    Location
	Description string
}

// Location - describes a place
type Location struct {
	ID          int
	Name        string
	Elevation   int
	Description string
	Longitude   float64
	Lattitude   float64
}

// StartSite - describes a start sight for flight
type StartSite struct {
	ID          int
	Name        string
	Description string
	Waypoints   []Waypoint
	Difficulty  int
}

// Waypoint - describes a start sight for flight
type Waypoint struct {
	ID         int
	Difficulty int
	Location   Location
}

// Flight - describes a flight
type Flight struct {
	ID          int
	User        User
	Startsite   StartSite
	Waypoint    Waypoint
	Duration    int
	Notes       string
	Distance    int
	MaxHight    int
	HangTime    int
	Wing        Wing
	Incidents   []Incident
	Photos      []FileReference
	FlightLog   FileReference
	FlightType  FlightType
	TakeOffType TakeoffType
}

// Incident - describes an incident
type Incident struct {
	ID          int
	Level       int
	Description string
	Public      bool
}

// Wing - describes a wing
type Wing struct {
	ID      int
	Name    string
	Images  []FileReference
	Details WingDetails
}

// WingDetails - details of the wing test
type WingDetails struct {
	ID          int
	Description string
	DhvScore    string
	EnaScore    string
}

// FileReference - describes whre a file is stored
type FileReference struct {
	MimeType     string
	FileName     string
	FileLocation string
}

// TakeoffType - describes a type of takeoff
type TakeoffType struct {
	ID    int
	Name  string
	Value int
}

// FlightType - is this a Paragliding, Speedrider, Baloon etc
type FlightType struct {
	ID    int
	Name  string
	Value int
}
