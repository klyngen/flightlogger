package common

// FlightLogDatabase - defines an interface to the data-layer
type FlightLogDatabase interface {
	// What all databases should do
	MigrateDatabase()
	CreateConnection(username string, password string, port string, hostname string) error

	// User CRUD
	CreateUser(user User) (User, error)
	GetAllUsers(limit int, page int) ([]User, error)
	GetUser(ID int) (User, error)
	UpdateUser(ID int, user User) (User, error)
	DeleteUser(ID int) error

	// Location CRUD and search
	CreateLocation(location Location) (Location, error)
	UpdateLocation(ID int, location Location) (Location, error)
	DeleteLocation(ID int) error
	LocationSearchByName(name string) ([]Location, error)
	GetLocation(ID int) (Location, error)

	// UserGroup CRUD and search
	CreateUserGroup(userGroup UserGroup, scopes []int) (UserGroup, error)
	UpdateUserGroup(groupID int, userGroup UserGroup, scopes []int) (UserGroup, error)
	GetAllUserGroups(limit int, page int) ([]UserGroup, error)
	GetUserGroup(ID int) (UserGroup, error)
	UserGroupSearchByName(name string) ([]UserGroup, error)

	// FileCreation CRD
	CreateFile(file FileReference) (FileReference, error)
	GetFile(ID int) (FileReference, error)
	DeleteFile(ID int) error

	// Flight CRUD
	CreateFlight(flight Flight) (Flight, error)
	UpdateFlight(ID int, flight Flight) (Flight, error)
	DeleteFlight(ID int) error
	GetAllFlights(limit int, page int) ([]Flight, error)
	GetFlight(ID int) (Flight, error)

	// FlightIncident CRUD and search
	CreateFlightIncident(incident Incident) (Incident, error)
	UpdateFlightIncident(ID int, Incident Incident) (Incident, error)
	DeleteFlightIncident(ID int) error
	GetFlightIncident(ID int) (Flight, error)
	GetFlightIncidentByLevel(errorLevel int) ([]Flight, error)
	GetFlightIncidents(limit int, page int) ([]Flight, error)

	// Wing CRUD
	CreateWing(wing Wing) (Wing, error)
	UpdateWing(ID int, wing Wing) (Wing, error)
	DeleteWing(ID int) error
	GetWing(ID int) (Wing, error)
	GetAllWings(limit int, page int) (Wing, error)
	GetWingSearchByName(name string) ([]Wing, error)

	// StartSite
	CreateStartSite(site StartSite) (StartSite, error)
	UpdateStartSite(ID int, site StartSite) (StartSite, error)
	DeleteStartSite(ID int) error
	GetStartStartSiteByName(name string) ([]StartSite, error)
	GetStartSiteByDifficulty(level int) ([]StartSite, error)
	GetStartSite(ID int) (StartSite, error)
	GetAllStartSites(limit int, page int) ([]StartSite, error)
}
