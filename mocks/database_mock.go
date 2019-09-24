package mocks

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/klyngen/flightlogger/common"
	"github.com/pkg/errors"
)

// RepoMock is made to mock a repository-implementation
type RepoMock struct{}

// What all databases should do
func (f *RepoMock) CreateConnection(username string, password string, database string, port string, hostname string) error {
	return nil
}

const (
	CRASH    string = "CRASH-CASE"
	PASSWORD string = "PASSWORD-CASE"
)

func (f *RepoMock) CreateUser(user *common.User) error {
	if user.Email == CRASH {
		return errors.New("Bad excuse for not creating a user")
	}

	user.ID = "GENERATED-ID"
	return nil
}

func createRandomUser() common.User {
	rand.NewSource(time.Now().UnixNano())
	return common.User{
		FirstName:    string(rand.Int()),
		LastName:     string(rand.Int()),
		Email:        fmt.Sprintf("crazy_%v@domain.com", rand.Int()),
		PasswordHash: []byte("HashyHash"),
		PasswordSalt: []byte("SaltySalt"),
	}
}

func (f *RepoMock) GetAllUsers(limit int, page int) ([]common.User, error) {
	users := []common.User{createRandomUser(), createRandomUser(), createRandomUser()}
	return users, nil
}

func (f *RepoMock) GetUser(ID string, user *common.User) error {
	if ID == CRASH {
		return errors.New("Crashing for no good reason... Deal with it")
	}

	usr := createRandomUser()

	if ID == PASSWORD {
		usr.PasswordHash = []byte("SOMETHING")
		usr.PasswordSalt = []byte("SOMETHING")
	}

	user = &usr

	return nil
}

func (f *RepoMock) GetUserByEmail(Email string, user *common.User) error {
	usr := createRandomUser()
	usr.Email = Email
	return nil
}

func (f *RepoMock) UpdateUser(ID string, user *common.User) error {
	if ID == CRASH {
		return errors.New("Crashing for no good reason... Deal with it")
	}

	return nil
}

func (f *RepoMock) DeleteUser(ID string) error {
	if ID == CRASH {
		return errors.New("Crashing for no good reason... Deal with it")
	}

	return nil
}

// Location CRUD and search
func (f *RepoMock) CreateLocation(location *common.Location) error {
	panic("not implemented")
}

func (f *RepoMock) UpdateLocation(ID uint, location *common.Location) error {
	panic("not implemented")
}

func (f *RepoMock) DeleteLocation(ID uint) error {
	panic("not implemented")
}

func (f *RepoMock) LocationSearchByName(name string, locations *common.Location) error {
	panic("not implemented")
}

func (f *RepoMock) GetLocation(ID uint, location *common.Location) error {
	panic("not implemented")
}

// UserGroup CRUD and search
func (f *RepoMock) CreateUserGroup(userGroup *common.UserGroup, scopes []int) error {
	panic("not implemented")
}

func (f *RepoMock) UpdateUserGroup(groupID uint, userGroup *common.UserGroup, permissions []int) error {
	panic("not implemented")
}

func (f *RepoMock) GetAllUserGroups(limit int, page int, userGroups []common.UserGroup) error {
	panic("not implemented")
}

func (f *RepoMock) GetUserGroup(ID uint, userGroup *common.UserGroup) error {
	panic("not implemented")
}

func (f *RepoMock) UserGroupSearchByName(name string, userGroups []common.UserGroup) error {
	panic("not implemented")
}

// FileCreation CRUD
func (f *RepoMock) CreateFile(file *common.FileReference) error {
	panic("not implemented")
}

func (f *RepoMock) GetFile(ID uint, file *common.FileReference) error {
	panic("not implemented")
}

func (f *RepoMock) DeleteFile(ID uint) error {
	panic("not implemented")
}

// Flight CRUD
func (f *RepoMock) CreateFlight(flight *common.Flight) error {
	panic("not implemented")
}

func (f *RepoMock) UpdateFlight(ID uint, flight *common.Flight) error {
	panic("not implemented")
}

func (f *RepoMock) DeleteFlight(ID uint, soft bool) error {
	panic("not implemented")
}

func (f *RepoMock) GetAllFlights(limit int, page int, flights []common.Flight) error {
	panic("not implemented")
}

func (f *RepoMock) GetFlight(ID uint, flight *common.Flight) error {
	panic("not implemented")
}

// FlightIncident CRUD and search
func (f *RepoMock) CreateFlightIncident(incident *common.Incident) error {
	panic("not implemented")
}

func (f *RepoMock) UpdateFlightIncident(ID uint, incident *common.Incident) error {
	panic("not implemented")
}

func (f *RepoMock) DeleteFlightIncident(ID uint) error {
	panic("not implemented")
}

func (f *RepoMock) GetFlightIncident(ID uint, flight *common.Flight) error {
	panic("not implemented")
}

func (f *RepoMock) GetFlightIncidentByLevel(errorLevel uint, flights []common.Flight) error {
	panic("not implemented")
}

func (f *RepoMock) GetFlightIncidents(limit int, page int, flights []common.Flight) error {
	panic("not implemented")
}

// Wing CRUD
func (f *RepoMock) CreateWing(wing *common.Wing) error {
	panic("not implemented")
}

func (f *RepoMock) UpdateWing(ID uint, wing *common.Wing) error {
	panic("not implemented")
}

func (f *RepoMock) DeleteWing(ID uint) error {
	panic("not implemented")
}

func (f *RepoMock) GetWing(ID uint, wing *common.Wing) (common.Wing, error) {
	panic("not implemented")
}

func (f *RepoMock) GetAllWings(limit uint, page uint, wing *common.Wing) error {
	panic("not implemented")
}

func (f *RepoMock) GetWingSearchByName(name string, wings []common.Wing) error {
	panic("not implemented")
}

// StartSite
func (f *RepoMock) CreateStartSite(site *common.StartSite) error {
	panic("not implemented")
}

func (f *RepoMock) UpdateStartSite(ID uint, site *common.StartSite) error {
	panic("not implemented")
}

func (f *RepoMock) DeleteStartSite(ID uint) error {
	panic("not implemented")
}

func (f *RepoMock) GetStartStartSiteByName(name string) ([]common.StartSite, error) {
	panic("not implemented")
}

func (f *RepoMock) GetStartSite(ID uint, startSite *common.StartSite) error {
	panic("not implemented")
}

func (f *RepoMock) GetAllStartSites(limit int, page int, startSites []common.StartSite) error {
	panic("not implemented")
}

func (f *RepoMock) GetSiteIncidents(siteID uint, incidents []common.Incident) error {
	panic("not implemented")
}

func (f *RepoMock) CreateWayPoint(point *common.Waypoint) error {
	panic("not implemented")
}

func (f *RepoMock) UpdateWayPoint(ID uint, point *common.Waypoint) error {
	panic("not implemented")
}
