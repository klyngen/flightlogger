package service

import "github.com/klyngen/flightlogger/common"

// FlightLogService describes our use-case layer

func (s *FlightLogService) UpdateUser(ID uint, user *common.User) error {
	panic("not implemented")
}

func (s *FlightLogService) DeleteUser(ID uint) error {
	panic("not implemented")
}

// Location CRUD and search
func (s *FlightLogService) CreateLocation(location *common.Location) error {
	panic("not implemented")
}

func (s *FlightLogService) UpdateLocation(ID uint, location *common.Location) error {
	panic("not implemented")
}

func (s *FlightLogService) DeleteLocation(ID uint) error {
	panic("not implemented")
}

func (s *FlightLogService) LocationSearchByName(name string) ([]common.Location, error) {
	panic("not implemented")
}

func (s *FlightLogService) GetLocation(ID uint, location *common.Location) error {
	panic("not implemented")
}

// UserGroup CRUD and search
func (s *FlightLogService) CreateUserGroup(userGroup *common.UserGroup, scopes []int) error {
	panic("not implemented")
}

func (s *FlightLogService) UpdateUserGroup(groupID uint, userGroup *common.UserGroup, scopes []int) error {
	panic("not implemented")
}

func (s *FlightLogService) GetAllUserGroups(limit int, page int) ([]common.UserGroup, error) {
	panic("not implemented")
}

func (s *FlightLogService) GetUserGroup(ID uint, userGroup *common.UserGroup) error {
	panic("not implemented")
}

func (s *FlightLogService) UserGroupSearchByName(name string) ([]common.UserGroup, error) {
	panic("not implemented")
}

// CreateFile stores the file somewhere and creates a record in the database
func (s *FlightLogService) CreateFile(file *common.FileReference) error {
	panic("not implemented")
}

func (s *FlightLogService) GetFile(ID uint, file *common.FileReference) error {
	panic("not implemented")
}

func (s *FlightLogService) DeleteFile(ID uint) error {
	panic("not implemented")
}

// Flight CRUD
func (s *FlightLogService) CreateFlight(flight *common.Flight) error {
	panic("not implemented")
}

func (s *FlightLogService) UpdateFlight(ID uint, flight *common.Flight) error {
	panic("not implemented")
}

func (s *FlightLogService) DeleteFlight(ID uint, soft bool) error {
	panic("not implemented")
}

func (s *FlightLogService) GetAllFlights(limit int, page int) ([]common.Flight, error) {
	panic("not implemented")
}

func (s *FlightLogService) GetFlight(ID uint, flight *common.Flight) error {
	panic("not implemented")
}

// FlightIncident CRUD and search
func (s *FlightLogService) CreateFlightIncident(incident *common.Incident) error {
	panic("not implemented")
}

func (s *FlightLogService) UpdateFlightIncident(ID uint, Incident *common.Incident) error {
	panic("not implemented")
}

func (s *FlightLogService) DeleteFlightIncident(ID uint) error {
	panic("not implemented")
}

func (s *FlightLogService) GetFlightIncident(ID uint, incident *common.Incident) error {
	panic("not implemented")
}

func (s *FlightLogService) GetFlightIncidentByLevel(errorLevel uint) ([]common.Incident, error) {
	panic("not implemented")
}

func (s *FlightLogService) GetFlightIncidents(limit int, page int) ([]common.Incident, error) {
	panic("not implemented")
}

// Wing CRUD
func (s *FlightLogService) CreateWing(wing *common.Wing) error {
	panic("not implemented")
}

func (s *FlightLogService) UpdateWing(ID uint, wing *common.Wing) error {
	panic("not implemented")
}

func (s *FlightLogService) DeleteWing(ID uint) error {
	panic("not implemented")
}

func (s *FlightLogService) GetWing(ID uint, wing *common.Wing) error {
	panic("not implemented")
}

func (s *FlightLogService) GetAllWings(limit uint, page uint) ([]common.Wing, error) {
	panic("not implemented")
}

func (s *FlightLogService) GetWingSearchByName(name string) ([]common.Wing, error) {
	panic("not implemented")
}

// StartSite
func (s *FlightLogService) CreateStartSite(site *common.StartSite) error {
	panic("not implemented")
}

func (s *FlightLogService) UpdateStartSite(ID uint, site *common.StartSite) error {
	panic("not implemented")
}

func (s *FlightLogService) DeleteStartSite(ID uint) error {
	panic("not implemented")
}

func (s *FlightLogService) GetStartStartSiteByName(name string) ([]common.StartSite, error) {
	panic("not implemented")
}

func (s *FlightLogService) GetStartSite(ID uint, startSite *common.StartSite) error {
	panic("not implemented")
}

func (s *FlightLogService) GetAllStartSites(limit int, page int) ([]common.StartSite, error) {
	panic("not implemented")
}

func (s *FlightLogService) GetSiteIncidents(siteID uint) ([]common.Incident, error) {
	panic("not implemented")
}

func (s *FlightLogService) CreateWayPoint(point *common.Waypoint) error {
	panic("not implemented")
}

func (s *FlightLogService) UpdateWayPoint(ID uint, point *common.Waypoint) error {
	panic("not implemented")
}
