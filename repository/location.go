package repository

import (
	"database/sql"
	"log"

	"github.com/klyngen/flightlogger/common"
)

// Basically create coordinates if they dont exist and return an ID anyway
func createCoordinates(db *sql.DB, longitude float64, lattitude float64) (uint, error) {
	// See if we have the coordinates already
	stmt, err := db.Prepare("SELECT Id FROM Coordinates WHERE Longitude = ? AND Latitude = ?")

	if err != nil {
		log.Printf("Badly formed query")
		return 0, err
	}

	var row uint

	err = stmt.QueryRow(longitude, lattitude).Scan(&row)

	stmt.Close()

	// Insert the coordinates into the database
	if err == sql.ErrNoRows {
		stmt, err = db.Prepare("INSERT INTO Coordinates (Longitude, Latitude) VALUES (?, ?)")

		defer stmt.Close()

		var res sql.Result
		if res, err = stmt.Exec(longitude, lattitude); err != nil {
			log.Printf("Could not insert the coordinate set: %v \n", err)
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
		log.Printf("Badly formed query for creating a country %v", err)
		return 0, err
	}

	var row uint

	err = stmt.QueryRow(name).Scan(&row)

	log.Println(row)

	stmt.Close()

	// Insert the coordinates into the database
	if err == sql.ErrNoRows {
		stmt, err = db.Prepare("INSERT INTO Country (Name) VALUES  (?)")

		defer stmt.Close()

		var res sql.Result
		if res, err = stmt.Exec(name); err != nil {
			log.Printf("Could not insert the country %v", err)
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

func createCountryPart(db *sql.DB, areaName string, postalCode string, countryName string) (uint, error) {

	countryID, err := createCountry(db, countryName)

	log.Println(countryID)

	if err != nil {
		log.Printf("Could not resolve all constraints for a countryPart %v", err)
		return 0, err
	}

	// See if we have the coordinates already
	stmt, err := db.Prepare("SELECT Id FROM CountryPart WHERE PostalCode = ? AND Areaname = ? AND CountryId = ?")

	if err != nil {
		log.Printf("Poorly formed query: %v", err)
		return 0, err
	}

	var row uint

	err = stmt.QueryRow(postalCode, areaName, countryID).Scan(&row)

	log.Println(row)

	stmt.Close()

	// Insert the coordinates into the database
	if err == sql.ErrNoRows {
		stmt, err = db.Prepare("INSERT INTO CountryPart (PostalCode, AreaName, CountryId) VALUES (?, ?, ?)")

		defer stmt.Close()

		var res sql.Result
		if res, err = stmt.Exec(postalCode, areaName, countryID); err != nil {
			log.Printf("Could not insert the country-part %v \n", err)
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
	if err != nil {
		log.Printf("Could not resolve location-keys %v", err)
		return err
	}
	countryPartID, err := createCountryPart(f.db, location.AreaName, location.PostalCode, location.CountryName)

	log.Println(coordinateID, countryPartID)

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

// UpdateLocation - updates location data and other directly connected entities
func (f *MySQLRepository) UpdateLocation(ID uint, location *common.Location) error {
	var stmt *sql.Stmt
	var err error
	if stmt, err = f.db.Prepare("UPDATE Location SET Name = ?, Elevation = ?, Description = ?, CountryPartId = ?, CoordinateId = ? WHERE Id = ?"); err != nil {
		log.Printf("Could not prepare update-location statement: %v \n", err)
		return err
	}
	defer stmt.Close()

	// Resolve ID's of sub-entities
	coordinateID, err := createCoordinates(f.db, location.Longitude, location.Lattitude)
	if err != nil {
		log.Printf("Could not resolve location-keys %v", err)
		return err
	}
	countryPartID, err := createCountryPart(f.db, location.AreaName, location.PostalCode, location.CountryName)

	if err != nil {
		log.Printf("Could not resolve keys %v", err)
		return err
	}

	var result sql.Result
	// Now we have our entities, we update!
	if result, err = stmt.Exec(location.Name, location.Elevation, location.Description, countryPartID, coordinateID, location.ID); err != nil {
		return err
	}

	var rows int64
	if rows, err = result.RowsAffected(); err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// DeleteLocation soft-deletes the location
func (f *MySQLRepository) DeleteLocation(ID uint) error {
	var stmt *sql.Stmt
	var err error
	if stmt, err = f.db.Prepare("UPDATE Location SET Deleted = CURRENT_TIMESTAMP WHERE Id = ? LIMIT 1"); err != nil {
		log.Printf("Could not prepare delete-location statement %v \n", err)
		return err
	}

	var result sql.Result
	if result, err = stmt.Exec(ID); err != nil {
		log.Printf("Something did not work while deleting the Location with Id %d,  %v \n", ID, err)
		return err
	}

	changes, err := result.RowsAffected()

	if err != nil {
		log.Printf("Could not get the resultset when treaing ID: %d,  %v \n", ID, err)
		return err
	}

	if changes == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// LocationSearchByName - gets all locations resembling the search-term
func (f *MySQLRepository) LocationSearchByName(name string) ([]common.Location, error) {
	var err error
	var stmt *sql.Stmt
	if stmt, err = f.db.Prepare(`select Id, Elevation, Description, Location.Name, Areaname, PostalCode, Country.Name, Longitude, Latitude from Location 
		INNER JOIN CountryPart ON CountryPart.Id = Location.CountryPartId 
		INNER JOIN Country ON CountryPart.CountryId = Country.Id 
		INNER JOIN Coordinates ON CoordinateId = Coordinates.Id
			WHERE Name Like '%?%'`); err != nil {
		log.Printf("Unable to prepare Location-search statement")
		return nil, err
	}

	rows, err := stmt.Query(name)

	if err != nil {
		log.Printf("Could not get locations %v when searching by name %s \n", err, name)
		return nil, err
	}

	locations := make([]common.Location, 0)

	for rows.Next() {
		location := common.Location{}
		if err = rows.Scan(&location.ID, &location.Elevation, &location.Description, &location.Name, &location.Name, &location.AreaName,
			&location.PostalCode, &location.CountryName, &location.Longitude, &location.Lattitude); err != nil {
			log.Printf("Could not serialize the row %v", err)
		}
		locations = append(locations, location)
	}

	return locations, nil
}

// GetLocation gets the location by it's ID
func (f *MySQLRepository) GetLocation(ID uint, location *common.Location) error {
	var err error
	var stmt *sql.Stmt
	if stmt, err = f.db.Prepare(`select Elevation, Description, Location.Name, Areaname, PostalCode, Country.Name, Longitude, Latitude from Location 
		INNER JOIN CountryPart ON CountryPart.Id = Location.CountryPartId 
		INNER JOIN Country ON CountryPart.CountryId = Country.Id 
		INNER JOIN Coordinates ON CoordinateId = Coordinates.Id
			WHERE Name Id = ?`); err != nil {
		log.Printf("Unable to prepare Location-search statement")
		return err
	}
	rows, err := stmt.Query(ID)

	if location == nil {
		location = &common.Location{}
	}

	return rows.Scan(&location.ID, &location.Elevation, &location.Description, &location.Name, &location.Name, &location.AreaName,
		&location.PostalCode, &location.CountryName, &location.Longitude, &location.Lattitude)
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
