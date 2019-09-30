package repository

import (
	"database/sql"
	"log"

	"github.com/klyngen/flightlogger/common"
)

// Basically create coordinates if they dont exist and return an ID anyway
func createCoordinates(db *sql.DB, longitude float64, lattitude float64) (uint, error) {
	// See if we have the coordinates already
	stmt, err := db.Prepare("SELECT Id FROM Coordinates WHERE Longitude = ? AND Lattitude = ?")

	if err != nil {
		log.Printf("Badly formed query")
		return 0, err
	}

	var row uint

	err = stmt.QueryRow(longitude, lattitude).Scan(&row)

	stmt.Close()

	// Insert the coordinates into the database
	if err == nil {
		stmt, err = db.Prepare("INSERT INTO Coordinates (Longitude, Latitude) VALUES (?, ?)")

		defer stmt.Close()

		var res sql.Result
		if res, err = stmt.Exec(longitude, lattitude); err != nil {
			log.Println("Could not insert the coordinate set")
			return 0, err
		}

		var ID uint
		if intres, err := res.LastInsertId(); err != nil {
			ID = uint(intres)
		}

		return ID, nil
	}

	return row, err
}

// Create the country if it does not exist
func createCountry(db *sql.DB, name string) (uint, error) {
	// See if we have the coordinates already
	stmt, err := db.Prepare("SELECT Id FROM Country WHERE Name = ?")

	if err != nil {
		log.Printf("Badly formed query")
		return 0, err
	}

	var row uint

	err = stmt.QueryRow(name).Scan(&row)

	stmt.Close()

	// Insert the coordinates into the database
	if err == nil {
		stmt, err = db.Prepare("INSERT INTO Country (Name) VALUES  (?)")

		defer stmt.Close()

		var res sql.Result
		if res, err = stmt.Exec(name); err != nil {
			log.Println("Could not insert the coordinate set")
			return 0, err
		}

		var ID uint
		if intres, err := res.LastInsertId(); err != nil {
			ID = uint(intres)
		}

		return ID, nil
	}

	return row, err
}

func createCountryPart(db *sql.DB, countryPart string, postalCode string, countryName string) (uint, error) {

	countryID, err := createCountry(db, countryName)

	if err != nil {
		log.Printf("Could not resolve all constraints %v", err)
		return 0, err
	}

	// See if we have the coordinates already
	stmt, err := db.Prepare("SELECT Id FROM CountryPart WHERE PostalCode = ? AND CountryPart = ? AND CountryId = ?")

	if err != nil {
		log.Printf("Poorly formed query: %v", err)
		return 0, err
	}

	var row uint

	err = stmt.QueryRow(postalCode, countryPart, countryID).Scan(&row)

	stmt.Close()

	// Insert the coordinates into the database
	if err == nil {
		stmt, err = db.Prepare("INSERT INTO CountryPart (PostalCode, AreaName, CountryId) VALUES (?, ?, ?)")

		defer stmt.Close()

		var res sql.Result
		if res, err = stmt.Exec(postalCode, countryPart, countryID); err != nil {
			log.Println("Could not insert the coordinate set")
			return 0, err
		}

		var ID uint
		if intres, err := res.LastInsertId(); err != nil {
			ID = uint(intres)
		}

		return ID, nil
	}

	return row, err
}

// CreateLocation creates a location with a set of coordinates if they dont exist
// a location is really the main entity and should only created together with its sub-entities
func (f *MySQLRepository) CreateLocation(location *common.Location) error {
	coordinateID, err := createCoordinates(f.db, location.Longitude, location.Lattitude)
	countryPartID, err := createCountryPart(f.db, location.CountryPart, location.PostalCode, location.CountryName)

	if err != nil {
		log.Printf("Could not resolve keys %v", err)
		return err
	}

	stmt, err := f.db.Prepare("INSERT INTO Location (CoordinateId, CountryPartId, Elevation, Description) VALUES (?, ?, ?, ?)")

	if err != nil {
		log.Printf("Could not prepare statement %v", err)
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(coordinateID, countryPartID, location.Elevation, location.Description)

	if err != nil {
		log.Printf("Could not insert the row %v", err)
		return err
	}

	rowID, err := res.LastInsertId()

	location.ID = uint(rowID)
	return err
}

func (f *MySQLRepository) UpdateLocation(ID uint, location *common.Location) error {
	panic("not implemented")
}

func (f *MySQLRepository) DeleteLocation(ID uint) error {
	panic("not implemented")
}

func (f *MySQLRepository) LocationSearchByName(name string, locations *common.Location) error {
	panic("not implemented")
}

func (f *MySQLRepository) GetLocation(ID uint, location *common.Location) error {
	panic("not implemented")
}

// StartSite
func (f *MySQLRepository) CreateStartSite(site *common.StartSite) error {
	panic("not implemented")
}

func (f *MySQLRepository) UpdateStartSite(ID uint, site *common.StartSite) error {
	panic("not implemented")
}

func (f *MySQLRepository) DeleteStartSite(ID uint) error {
	panic("not implemented")
}

func (f *MySQLRepository) GetStartStartSiteByName(name string) ([]common.StartSite, error) {
	panic("not implemented")
}

func (f *MySQLRepository) GetStartSite(ID uint, startSite *common.StartSite) error {
	panic("not implemented")
}

func (f *MySQLRepository) GetAllStartSites(limit int, page int, startSites []common.StartSite) error {
	panic("not implemented")
}

func (f *MySQLRepository) GetSiteIncidents(siteID uint, incidents []common.Incident) error {
	panic("not implemented")
}

func (f *MySQLRepository) CreateWayPoint(point *common.Waypoint) error {
	panic("not implemented")
}

func (f *MySQLRepository) UpdateWayPoint(ID uint, point *common.Waypoint) error {
	panic("not implemented")
}
