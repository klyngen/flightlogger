package common

// FlightLogService describes how we interact with our business-logic
type FlightLogService interface {
	Authenticate(user *User) error

	CreateUser(user User) (User, error)
	GetAllUsers(limit int, page int) ([]User, error)
	GetUser(ID uint) (User, error)
	UpdateUser(ID uint, user User) (User, error)
	DeleteUser(ID uint) error

	// Location CRUD and search
	CreateLocation(location Location) (Location, error)
	UpdateLocation(ID uint, location Location) (Location, error)
	DeleteLocation(ID uint) error
	LocationSearchByName(name string) ([]Location, error)
	GetLocation(ID uint) (Location, error)

	// UserGroup CRUD and search
	CreateUserGroup(userGroup UserGroup, scopes []int) (UserGroup, error)
	UpdateUserGroup(groupID uint, userGroup UserGroup, scopes []int) (UserGroup, error)
	GetAllUserGroups(limit int, page int) ([]UserGroup, error)
	GetUserGroup(ID uint) (UserGroup, error)
	UserGroupSearchByName(name string) ([]UserGroup, error)

	// FileCreation CRD
	CreateFile(file FileReference) (FileReference, error)
	GetFile(ID uint) (FileReference, error)
	DeleteFile(ID uint) error

	// Flight CRUD
	CreateFlight(flight Flight) (Flight, error)
	UpdateFlight(ID uint, flight Flight) (Flight, error)
	DeleteFlight(ID uint, soft bool) error
	GetAllFlights(limit int, page int) ([]Flight, error)
	GetFlight(ID uint) (Flight, error)

	// FlightIncident CRUD and search
	CreateFlightIncident(incident Incident) (Incident, error)
	UpdateFlightIncident(ID uint, Incident Incident) (Incident, error)
	DeleteFlightIncident(ID uint) error
	GetFlightIncident(ID uint) (Flight, error)
	GetFlightIncidentByLevel(errorLevel uint) ([]Flight, error)
	GetFlightIncidents(limit int, page int) ([]Flight, error)

	// Wing CRUD
	CreateWing(wing Wing) (Wing, error)
	UpdateWing(ID uint, wing Wing) (Wing, error)
	DeleteWing(ID uint) error
	GetWing(ID uint) (Wing, error)
	GetAllWings(limit uint, page uint) (Wing, error)
	GetWingSearchByName(name string) ([]Wing, error)

	// StartSite
	CreateStartSite(site StartSite) (StartSite, error)
	UpdateStartSite(ID uint, site StartSite) (StartSite, error)
	DeleteStartSite(ID uint) error
	GetStartStartSiteByName(name string) ([]StartSite, error)
	GetStartSite(ID uint) (StartSite, error)
	GetAllStartSites(limit int, page int) ([]StartSite, error)
	GetSiteIncidents(siteID uint) ([]Incident, error)

	CreateWayPoint(point Waypoint) (Waypoint, error)
	UpdateWayPoint(ID uint, point Waypoint) (Waypoint, error)
}
