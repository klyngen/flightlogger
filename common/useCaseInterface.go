package common

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/casbin/casbin"
	"github.com/dgrijalva/jwt-go"
)

// FlightLogService describes how we interact with our business-logic
type FlightLogService interface {
	GetSessionManager() *scs.SessionManager
	GetCasbinEnforcer() *casbin.Enforcer
	Authenticate(username string, password string, request *http.Request) error
	VerifyTokenString(token string) (jwt.Claims, error)
	ActivateUser(UserID string) error

	CreateUser(user *User, password string) error
	GetAllUsers(limit int, page int) ([]User, error)
	GetUser(ID string, user *User) error
	UpdateUser(ID uint, user *User) error
	DeleteUser(ID uint) error

	// Location CRUD and search
	CreateLocation(location *Location) error
	UpdateLocation(ID uint, location *Location) error
	DeleteLocation(ID uint) error
	LocationSearchByName(name string) ([]Location, error)
	GetLocation(ID uint, location *Location) error

	// FileCreation CRUD
	CreateFile(file *FileReference) error
	GetFile(ID uint, file *FileReference) error
	DeleteFile(ID uint) error

	// Flight CRUD
	CreateFlight(flight *Flight) error
	UpdateFlight(ID uint, flight *Flight) error
	DeleteFlight(ID uint, soft bool) error
	GetAllFlights(limit int, page int) ([]Flight, error)
	GetFlight(ID uint, flight *Flight) error

	// FlightIncident CRUD and search
	CreateFlightIncident(incident *Incident) error
	UpdateFlightIncident(ID uint, incident *Incident) error
	DeleteFlightIncident(ID uint) error
	GetFlightIncident(ID uint, flight *Incident) error
	GetFlightIncidentByLevel(errorLevel uint) ([]Incident, error)
	GetFlightIncidents(limit int, page int) ([]Incident, error)

	// Wing CRUD
	CreateWing(wing *FlyingDevice) error
	UpdateWing(ID uint, wing *FlyingDevice) error
	DeleteWing(ID uint) error
	GetWing(ID uint, wing *FlyingDevice) error
	GetAllWings(limit uint, page uint) ([]FlyingDevice, error)
	GetWingSearchByName(name string) ([]FlyingDevice, error)

	// StartSite
	CreateStartSite(site *StartSite) error
	UpdateStartSite(ID uint, site *StartSite) error
	DeleteStartSite(ID uint) error
	GetStartStartSiteByName(name string) ([]StartSite, error)
	GetStartSite(ID uint, startSite *StartSite) error
	GetAllStartSites(limit int, page int) ([]StartSite, error)
	GetSiteIncidents(siteID uint) ([]Incident, error)

	CreateWayPoint(point *Waypoint) error
	UpdateWayPoint(ID uint, point *Waypoint) error
}
