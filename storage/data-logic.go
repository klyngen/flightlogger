package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/klyngen/flightlogger/common"
)

// OrmDatabase - should implement the databaseInterface
type OrmDatabase struct {
	db *gorm.DB
}

// MigrateDatabase - migrates the database
func (d *OrmDatabase) MigrateDatabase() {
	// Migrate location first
	d.db.AutoMigrate(&DbFileReference{})
	d.db.AutoMigrate(&DbCoordinates{})
	d.db.AutoMigrate(&DbLocation{})

	// Create club entity before user and flights
	d.db.AutoMigrate(&DbClub{})

	// Waypoint and start are dependent on location
	d.db.AutoMigrate(&DbWaypoint{})
	d.db.AutoMigrate(&DbStartSite{})

	// Wing related data
	d.db.AutoMigrate(&DbWingScoreDetails{})
	d.db.AutoMigrate(&DbWing{})

	// Flight related entities
	d.db.AutoMigrate(&DbFlightType{})
	d.db.AutoMigrate(&DbTakeoffType{})
	d.db.AutoMigrate(&DbIncident{})
	d.db.AutoMigrate(&DbFlight{})

	// Set up the user related entities
	d.db.AutoMigrate(&DbCredentials{})
	d.db.AutoMigrate(&DbUserScope{})
	d.db.AutoMigrate(&DbUserGroup{})
	d.db.AutoMigrate(&DbUser{})
}

// CreateConnection - establish a connection to the database
func (d *OrmDatabase) CreateConnection(username string, password string, database string, port string, hostname string) error {
	panic("not implemented")
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%?charset=utf8&parseTime=True&loc=%s", username, password, database, hostname))

	if err != nil {
		return err
	}

	d.db = db
	return nil
}

// ############## USER RELATED QUERIES ############################

// CreateUser - try to create a new user
func (d *OrmDatabase) CreateUser(user common.User) (common.User, error) {
	mappedUser := mapUser(user)
	err := d.db.Create(&mappedUser).Error

	if err != nil {
		user.ID = mappedUser.Model.ID
		return user, err
	}

	return user, nil
}

// GetAllUsers - gets all users
func (d *OrmDatabase) GetAllUsers(limit int, page int) ([]common.User, error) {
	var users []DbUser
	d.db.Limit(limit).Offset((page - 1) * limit).Find(&users)

}

func (d *OrmDatabase) GetUser(ID int) (common.User, error) {
	panic("not implemented")
}

func (d *OrmDatabase) UpdateUser(ID int, user common.User) (common.User, error) {
	panic("not implemented")
}

func (d *OrmDatabase) DeleteUser(ID int) error {
	panic("not implemented")
}

// Location CRUD and search
func (d *OrmDatabase) CreateLocation(location common.Location) (common.Location, error) {
	panic("not implemented")
}

func (d *OrmDatabase) UpdateLocation(ID int, location common.Location) (common.Location, error) {
	panic("not implemented")
}

func (d *OrmDatabase) DeleteLocation(ID int) error {
	panic("not implemented")
}

func (d *OrmDatabase) LocationSearchByName(name string) ([]common.Location, error) {
	panic("not implemented")
}

func (d *OrmDatabase) GetLocation(ID int) (common.Location, error) {
	panic("not implemented")
}

// UserGroup CRUD and search
func (d *OrmDatabase) CreateUserGroup(userGroup common.UserGroup, scopes []int) (common.UserGroup, error) {
	panic("not implemented")
}

func (d *OrmDatabase) UpdateUserGroup(groupID int, userGroup common.UserGroup, scopes []int) (common.UserGroup, error) {
	panic("not implemented")
}

func (d *OrmDatabase) GetAllUserGroups(limit int) ([]common.UserGroup, error) {
	panic("not implemented")
}

func (d *OrmDatabase) GetUserGroup(ID int) (common.UserGroup, error) {
	panic("not implemented")
}

func (d *OrmDatabase) UserGroupSearchByName(name string) ([]common.UserGroup, error) {
	panic("not implemented")
}

// FileCreation CRD
func (d *OrmDatabase) CreateFile(file common.FileReference) (common.FileReference, error) {
	panic("not implemented")
}

func (d *OrmDatabase) GetFile(ID int) (common.FileReference, error) {
	panic("not implemented")
}

func (d *OrmDatabase) DeleteFile(ID int) error {
	panic("not implemented")
}

// Flight CRUD
func (d *OrmDatabase) CreateFlight(flight common.Flight) (common.Flight, error) {
	panic("not implemented")
}

func (d *OrmDatabase) UpdateFlight(ID int, flight common.Flight) (common.Flight, error) {
	panic("not implemented")
}

func (d *OrmDatabase) DeleteFlight(ID int) error {
	panic("not implemented")
}

func (d *OrmDatabase) GetAllFlights(limit int) ([]common.Flight, error) {
	panic("not implemented")
}

func (d *OrmDatabase) GetFlight(ID int) (common.Flight, error) {
	panic("not implemented")
}

// FlightIncident CRUD and search
func (d *OrmDatabase) CreateFlightIncident(incident common.Incident) (common.Incident, error) {
	panic("not implemented")
}

func (d *OrmDatabase) UpdateFlightIncident(ID int, Incident common.Incident) (common.Incident, error) {
	panic("not implemented")
}

func (d *OrmDatabase) DeleteFlightIncident(ID int) error {
	panic("not implemented")
}

func (d *OrmDatabase) GetFlightIncident(ID int) (common.Flight, error) {
	panic("not implemented")
}

func (d *OrmDatabase) GetFlightIncidentByLevel(errorLevel int) ([]common.Flight, error) {
	panic("not implemented")
}

func (d *OrmDatabase) GetFlightIncidents(limit int) ([]common.Flight, error) {
	panic("not implemented")
}

// Wing CRUD
func (d *OrmDatabase) CreateWing(wing common.Wing) (common.Wing, error) {
	panic("not implemented")
}

func (d *OrmDatabase) UpdateWing(ID int, wing common.Wing) (common.Wing, error) {
	panic("not implemented")
}

func (d *OrmDatabase) DeleteWing(ID int) error {
	panic("not implemented")
}

func (d *OrmDatabase) GetWing(ID int) (common.Wing, error) {
	panic("not implemented")
}

func (d *OrmDatabase) GetAllWings(limit int) (common.Wing, error) {
	panic("not implemented")
}

func (d *OrmDatabase) GetWingSearchByName(name string) ([]common.Wing, error) {
	panic("not implemented")
}

// StartSite
func (d *OrmDatabase) CreateStartSite(site common.StartSite) (common.StartSite, error) {
	panic("not implemented")
}

func (d *OrmDatabase) UpdateStartSite(ID int, site common.StartSite) (common.StartSite, error) {
	panic("not implemented")
}

func (d *OrmDatabase) DeleteStartSite(ID int) error {
	panic("not implemented")
}

func (d *OrmDatabase) GetStartStartSiteByName(name string) ([]common.StartSite, error) {
	panic("not implemented")
}

func (d *OrmDatabase) GetStartSiteByDifficulty(level int) ([]common.StartSite, error) {
	panic("not implemented")
}

func (d *OrmDatabase) GetStartSite(ID int) (common.StartSite, error) {
	panic("not implemented")
}

func (d *OrmDatabase) GetAllStartSites(limit int) ([]common.StartSite, error) {
	panic("not implemented")
}
