package usecase

import "github.com/klyngen/flightlogger/common"

// Service describes our use-case layer
type Service struct {
	database common.FlightLogDatabase
}

func NewService(database common.FlightLogDatabase) *Service {
	return &Service{database: database}
}

func (s *Service) Authenticate(user *common.User) error {
	panic("not implemented")
}

func (s *Service) CreateUser(user common.User) (common.User, error) {
	return s.database.CreateUser(user)
}

func (s *Service) GetAllUsers(limit int, page int) ([]common.User, error) {
	return s.database.GetAllUsers(limit, page)
}

func (s *Service) GetUser(ID uint) (common.User, error) {
	return s.database.GetUser(ID)
}

func (s *Service) UpdateUser(ID uint, user common.User) (common.User, error) {
	panic("not implemented")
}

func (s *Service) DeleteUser(ID uint) error {
	panic("not implemented")
}

// Location CRUD and search
func (s *Service) CreateLocation(location common.Location) (common.Location, error) {
	panic("not implemented")
}

func (s *Service) UpdateLocation(ID uint, location common.Location) (common.Location, error) {
	panic("not implemented")
}

func (s *Service) DeleteLocation(ID uint) error {
	panic("not implemented")
}

func (s *Service) LocationSearchByName(name string) ([]common.Location, error) {
	panic("not implemented")
}

func (s *Service) GetLocation(ID uint) (common.Location, error) {
	panic("not implemented")
}

// UserGroup CRUD and search
func (s *Service) CreateUserGroup(userGroup common.UserGroup, scopes []int) (common.UserGroup, error) {
	panic("not implemented")
}

func (s *Service) UpdateUserGroup(groupID uint, userGroup common.UserGroup, scopes []int) (common.UserGroup, error) {
	panic("not implemented")
}

func (s *Service) GetAllUserGroups(limit int, page int) ([]common.UserGroup, error) {
	panic("not implemented")
}

func (s *Service) GetUserGroup(ID uint) (common.UserGroup, error) {
	panic("not implemented")
}

func (s *Service) UserGroupSearchByName(name string) ([]common.UserGroup, error) {
	panic("not implemented")
}

// FileCreation CRD
func (s *Service) CreateFile(file common.FileReference) (common.FileReference, error) {
	panic("not implemented")
}

func (s *Service) GetFile(ID uint) (common.FileReference, error) {
	panic("not implemented")
}

func (s *Service) DeleteFile(ID uint) error {
	panic("not implemented")
}

// Flight CRUD
func (s *Service) CreateFlight(flight common.Flight) (common.Flight, error) {
	panic("not implemented")
}

func (s *Service) UpdateFlight(ID uint, flight common.Flight) (common.Flight, error) {
	panic("not implemented")
}

func (s *Service) DeleteFlight(ID uint, soft bool) error {
	panic("not implemented")
}

func (s *Service) GetAllFlights(limit int, page int) ([]common.Flight, error) {
	panic("not implemented")
}

func (s *Service) GetFlight(ID uint) (common.Flight, error) {
	panic("not implemented")
}

// FlightIncident CRUD and search
func (s *Service) CreateFlightIncident(incident common.Incident) (common.Incident, error) {
	panic("not implemented")
}

func (s *Service) UpdateFlightIncident(ID uint, Incident common.Incident) (common.Incident, error) {
	panic("not implemented")
}

func (s *Service) DeleteFlightIncident(ID uint) error {
	panic("not implemented")
}

func (s *Service) GetFlightIncident(ID uint) (common.Flight, error) {
	panic("not implemented")
}

func (s *Service) GetFlightIncidentByLevel(errorLevel uint) ([]common.Flight, error) {
	panic("not implemented")
}

func (s *Service) GetFlightIncidents(limit int, page int) ([]common.Flight, error) {
	panic("not implemented")
}

// Wing CRUD
func (s *Service) CreateWing(wing common.Wing) (common.Wing, error) {
	panic("not implemented")
}

func (s *Service) UpdateWing(ID uint, wing common.Wing) (common.Wing, error) {
	panic("not implemented")
}

func (s *Service) DeleteWing(ID uint) error {
	panic("not implemented")
}

func (s *Service) GetWing(ID uint) (common.Wing, error) {
	panic("not implemented")
}

func (s *Service) GetAllWings(limit uint, page uint) (common.Wing, error) {
	panic("not implemented")
}

func (s *Service) GetWingSearchByName(name string) ([]common.Wing, error) {
	panic("not implemented")
}

// StartSite
func (s *Service) CreateStartSite(site common.StartSite) (common.StartSite, error) {
	panic("not implemented")
}

func (s *Service) UpdateStartSite(ID uint, site common.StartSite) (common.StartSite, error) {
	panic("not implemented")
}

func (s *Service) DeleteStartSite(ID uint) error {
	panic("not implemented")
}

func (s *Service) GetStartStartSiteByName(name string) ([]common.StartSite, error) {
	panic("not implemented")
}

func (s *Service) GetStartSite(ID uint) (common.StartSite, error) {
	panic("not implemented")
}

func (s *Service) GetAllStartSites(limit int, page int) ([]common.StartSite, error) {
	panic("not implemented")
}

func (s *Service) GetSiteIncidents(siteID uint) ([]common.Incident, error) {
	panic("not implemented")
}

func (s *Service) CreateWayPoint(point common.Waypoint) (common.Waypoint, error) {
	panic("not implemented")
}

func (s *Service) UpdateWayPoint(ID uint, point common.Waypoint) (common.Waypoint, error) {
	panic("not implemented")
}
