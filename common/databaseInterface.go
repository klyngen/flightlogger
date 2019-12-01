package common

// FlightLogDatabase - defines an interface to the data-layer
type FlightLogDatabase interface {
	// What all databases should do
	CreateConnection(username string, password string, database string, port string, hostname string) error

	CreateUser(user *User) error
	GetAllUsers(limit int, page int) ([]User, error)
	GetUser(ID string, user *User) error
	GetUserByEmail(Email string, user *User) error
	UpdateUser(ID string, user *User) error
	DeleteUser(ID string) error

	ActivateUser(UserID string) error

	// Location CRUD and search
	CreateLocation(location *Location) error
	UpdateLocation(ID uint, location *Location) error
	DeleteLocation(ID uint) error
	LocationSearchByName(name string) ([]Location, error)
	GetLocation(ID uint, location *Location) error

	// UserGroup CRUD and search
	CreateUserGroup(userGroup *UserGroup, scopes []int) error
	UpdateUserGroup(groupID uint, userGroup *UserGroup, permissions []int) error
	GetAllUserGroups(limit int, page int, userGroups []UserGroup) error
	GetUserGroup(ID uint, userGroup *UserGroup) error
	UserGroupSearchByName(name string, userGroups []UserGroup) error

	// FileCreation CRUD
	CreateFile(file *FileReference) error
	GetFile(ID uint, file *FileReference) error
	DeleteFile(ID uint) error

	// Flight CRUD
	CreateFlight(flight *Flight) error
	UpdateFlight(ID uint, flight *Flight) error
	DeleteFlight(ID uint, soft bool) error
	GetAllFlights(limit int, page int, flights []Flight) error
	GetFlight(ID uint, flight *Flight) error

	// FlightIncident CRUD and search
	CreateFlightIncident(incident *Incident) error
	UpdateFlightIncident(ID uint, incident *Incident) error
	DeleteFlightIncident(ID uint) error
	GetFlightIncident(ID uint, flight *Flight) error
	GetFlightIncidentByLevel(errorLevel uint, flights []Flight) error
	GetFlightIncidents(limit int, page int, flights []Flight) error

	// Wing CRUD
	CreateWing(wing *FlyingDevice) error
	UpdateWing(ID uint, wing *FlyingDevice) error
	DeleteWing(ID uint) error
	GetWing(ID uint, wing *FlyingDevice) error
	GetAllWings(limit uint, page uint) ([]FlyingDevice, error)
	GetWingSearchByName(name string, wings []FlyingDevice) error

	// StartSite
	CreateStartSite(site *StartSite) error
	UpdateStartSite(ID uint, site *StartSite) error
	DeleteStartSite(ID uint) error
	GetStartSite(ID uint, startSite *StartSite) error
	GetAllStartSites(limit int, page int) ([]StartSite, error)
	GetSiteIncidents(siteID uint, incidents []Incident) error
	GetStartSiteWaypoints(siteID uint) ([]Waypoint, error)

	CreateWayPoint(point *Waypoint) error
	DeleteWaypoint(ID uint) error
	UpdateWayPoint(ID uint, point *Waypoint) error
	GetWaypoint(ID uint, point *Waypoint) error
	GetAllWaypoints(limit int, page int) ([]Waypoint, error)
}
