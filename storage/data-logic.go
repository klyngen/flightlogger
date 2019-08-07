package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// This import is needed in order to utilize MySql
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/klyngen/flightlogger/common"
	"github.com/pkg/errors"
)

// OrmDatabase - should implement the databaseInterface
type OrmDatabase struct {
	db *gorm.DB
}

// MigrateDatabase - migrates the database
func (d *OrmDatabase) MigrateDatabase() error {
	// Migrate location first
	err := d.db.AutoMigrate(&DbFileReference{}).Error
	err = d.db.AutoMigrate(&DbCoordinates{}).Error
	err = d.db.AutoMigrate(&DbLocation{}).Error

	if err != nil {
		return errors.Wrap(err, "Unable to migrate basic Location-coordinates")
	}

	// Create club entity before user and flights
	err = d.db.AutoMigrate(&DbClub{}).Error

	// Waypoint and start are dependent on location
	err = d.db.AutoMigrate(&DbWaypoint{}).Error
	err = d.db.AutoMigrate(&DbStartSite{}).Error

	if err != nil {
		return errors.Wrap(err, "Unable to migrate flight base-entities")
	}

	// Wing related data
	err = d.db.AutoMigrate(&DbWingScoreDetails{}).Error
	err = d.db.AutoMigrate(&DbWing{}).Error

	if err != nil {
		return errors.Wrap(err, "Unable to migrate wing-entities")
	}

	// Flight related entities
	err = d.db.AutoMigrate(&DbFlightType{}).Error
	err = d.db.AutoMigrate(&DbTakeoffType{}).Error
	err = d.db.AutoMigrate(&DbIncident{}).Error
	err = d.db.AutoMigrate(&DbFlight{}).Error

	if err != nil {
		return errors.Wrap(err, "Unable to migrate flight-entities")
	}

	// Set up the user related entities
	err = d.db.AutoMigrate(&DbCredentials{}).Error
	err = d.db.AutoMigrate(&DbUserScope{}).Error
	err = d.db.AutoMigrate(&DbUserGroup{}).Error
	err = d.db.AutoMigrate(&DbUser{}).Error

	return errors.Wrap(err, "Unable to migrate the database")
}

// CreateConnection - establish a connection to the database
func (d *OrmDatabase) CreateConnection(username string, password string, database string, port string, hostname string) error {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, database))

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
	return demapUsers(users), nil
}

// GetUser - gets a single user if it exists
func (d *OrmDatabase) GetUser(ID int) (common.User, error) {
	var user DbUser
	return demapUser(user), errors.Wrap(d.db.First(user, ID).Error, "Unable to get user")
}

// UpdateUser - update an existing user if it exists
func (d *OrmDatabase) UpdateUser(ID int, user common.User) (common.User, error) {
	dbUser := mapUser(user)
	return demapUser(dbUser), errors.Wrap(d.db.Save(&dbUser).Error, "Unable to update a user")
}

// DeleteUser - deletes a user
func (d *OrmDatabase) DeleteUser(ID int) error {
	var user DbUser

	// FIXME: see if this actually works

	// TODO: also delete related entities #GDPR

	err := d.db.First(&user, ID).Error

	if err != nil {
		return errors.Wrap(err, "Cannot delete a user we cannot find")
	}

	err = d.db.Delete(&user, ID).Error

	if err != nil {
		return errors.Wrap(err, "Unable to delete the user")
	}

	return nil
}

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

func (d *OrmDatabase) GetAllUserGroups(limit int, page int) ([]common.UserGroup, error) {
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

func (d *OrmDatabase) GetAllFlights(limit int, page int) ([]common.Flight, error) {
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

func (d *OrmDatabase) GetFlightIncidents(limit int, page int) ([]common.Flight, error) {
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

func (d *OrmDatabase) GetAllWings(limit int, page int) (common.Wing, error) {
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

func (d *OrmDatabase) GetAllStartSites(limit int, page int) ([]common.StartSite, error) {
	panic("not implemented")
}

func (d *OrmDatabase) GetSiteIncidents(siteID uint) ([]common.Incident, error) {
	panic("not implemented")
}
