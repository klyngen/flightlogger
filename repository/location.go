package repository

import (
	"database/sql"
	"fmt"
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

	if err != nil {
		log.Printf("Could not resolve keys %v", err)
		return err
	}

	stmt, err := f.db.Prepare("INSERT INTO Location (Name, CoordinateId, CountryPartId, Elevation, Description) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		log.Printf("Could not prepare statement %v", err)
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(location.Name, coordinateID, countryPartID, location.Elevation, location.Description)

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
	if stmt, err = f.db.Prepare("UPDATE Location SET DeletedTimestamp = CURRENT_TIMESTAMP WHERE Id = ? LIMIT 1"); err != nil {
		log.Printf("Could not prepare delete-location statement %v \n", err)
		return err
	}

	var result sql.Result
	if result, err = stmt.Exec(ID); err != nil {
		log.Printf("Something did not work while deleting the Location with Id %d,  %v \n", ID, err)
		return err
	}
	defer stmt.Close()

	changes, err := result.RowsAffected()

	if err != nil {
		log.Printf("Could not get the resultset when ID: %d,  %v \n", ID, err)
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
	if stmt, err = f.db.Prepare("SELECT * FROM getlocation WHERE LocationName LIKE ?"); err != nil {
		log.Printf("Unable to prepare Location-search statement")
		return nil, err
	}

	rows, err := stmt.Query(fmt.Sprintf("%%%v%%", name))

	if err != nil {
		log.Printf("Could not get locations %v when searching by name %s \n", err, name)
		return nil, err
	}
	defer stmt.Close()

	locations := make([]common.Location, 0)

	for rows.Next() {
		location := common.Location{}
		if err = rows.Scan(&location.ID, &location.Elevation, &location.Description, &location.Name, &location.AreaName,
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
	if stmt, err = f.db.Prepare("SELECT * FROM getlocation WHERE LocationId = ?"); err != nil {
		log.Printf("Unable to prepare Location-search statement")
		return err
	}
	defer stmt.Close()

	return mapLocation(location, stmt.QueryRow(ID))
}

func mapLocation(location *common.Location, row *sql.Row) error {
	if location == nil {
		location = &common.Location{}
	}

	return row.Scan(&location.ID, &location.Elevation, &location.Description, &location.Name, &location.AreaName,
		&location.PostalCode, &location.CountryName, &location.Longitude, &location.Lattitude)
}

// CreateStartSite creates a new startsite connected to a Location
func (f *MySQLRepository) CreateStartSite(site *common.StartSite) error {
	var stmt *sql.Stmt
	var err error

	// Make a statement!
	if stmt, err = f.db.Prepare("INSERT INTO Startsite (Name, LocationId, Description, Difficulty) VALUES (?, ?, ?, ?)"); err != nil {
		log.Printf("Unable to prepare statement")
		return err
	}
	defer stmt.Close()

	var result sql.Result
	if result, err = stmt.Exec(site.Name, site.Location.ID, site.Description, site.Difficulty); err != nil {
		log.Printf("Unable to create the startsite %v", err)
		return err
	}

	var siteid int64
	if siteid, err = result.LastInsertId(); err != nil {
		log.Printf("Cannot get the ID of the startsite %v", err)
		return err
	}

	site.ID = uint(siteid)

	return nil
}

// UpdateStartSite updates a startsite
func (f *MySQLRepository) UpdateStartSite(ID uint, site *common.StartSite) error {
	var stmt *sql.Stmt
	var err error

	// Make a statement!
	if stmt, err = f.db.Prepare("UPDATE Startsite SET Name = ?, LocationId = ?, Description = ?, Difficulty = ? WHERE Id = ?"); err != nil {
		log.Printf("Unable to prepare statement")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(site.Name, site.Location.ID, site.Description, site.Difficulty, site.ID)

	return err
}

// DeleteStartSite soft-deletes a startsite
func (f *MySQLRepository) DeleteStartSite(ID uint) error {
	var stmt *sql.Stmt
	var err error

	// Make a statement!
	if stmt, err = f.db.Prepare("UPDATE Startsite SET Deleted = CURRENT_TIMESTAMP WHERE Id = ?"); err != nil {
		log.Printf("Unable to prepare delete start-site query %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ID)

	return err
}

// GetStartSite get a startsite using Location and StartSite view's
func (f *MySQLRepository) GetStartSite(ID uint, startSite *common.StartSite) error {
	var err error
	var stmt *sql.Stmt
	if stmt, err = f.db.Prepare("SELECT * FROM getstartsite WHERE StartsiteId = ?"); err != nil {
		log.Printf("Unable to prepare Location-search statement")
		return err
	}
	defer stmt.Close()

	return mapStartSite(stmt.QueryRow(ID), startSite)
}

func mapStartSite(row rowScanner, startsite *common.StartSite) error {
	if startsite == nil {
		startsite = &common.StartSite{}
		startsite.Location = common.Location{}
		startsite.Waypoints = make([]common.Waypoint, 0)
	}

	var tempWaypointId interface{}

	err := row.Scan(&startsite.Location.ID, &startsite.Location.Elevation, &startsite.Location.Description, &startsite.Location.Name, &startsite.Location.AreaName, &startsite.Location.PostalCode, &startsite.Location.CountryName, &startsite.Location.Longitude, &startsite.Location.Lattitude, &startsite.ID, &startsite.Name, &startsite.Description, &startsite.Difficulty, &tempWaypointId)

	// We decide whether we want to include the Waypoint
	if err != nil {
		return err
	}

	if tempWaypointId != "" {
		if tempId, ok := tempWaypointId.(uint); ok {
			startsite.Waypoints = append(startsite.Waypoints, common.Waypoint{ID: tempId})
		}
	}

	return nil
}

// GetAllStartSites is a paged fetch for start-sites
func (f *MySQLRepository) GetAllStartSites(limit int, page int) ([]common.StartSite, error) {
	var stmt *sql.Stmt
	var err error

	if stmt, err = f.db.Prepare("SELECT * FROM getstartsite LIMIT ?,?"); err != nil {
		log.Printf("Unable to prepare statement for retreiving all startsites %v", err)
		return nil, err
	}

	rows, err := stmt.Query((page-1)*limit, limit)
	defer stmt.Close()

	if err != nil {
		log.Printf("Could not get any results %v", err)
		return nil, err
	}

	defer rows.Close()

	sites := make([]common.StartSite, 0)

	for rows.Next() {
		var site common.StartSite

		mapStartSite(rows, &site)

		sites = append(sites, site)
	}
	return sites, nil
}

func (f *MySQLRepository) GetSiteIncidents(siteID uint, incidents []common.Incident) error {
	panic("not implemented")
}

// CreateWayPoint creates a waypoint
func (f *MySQLRepository) CreateWayPoint(point *common.Waypoint) error {
	var stmt *sql.Stmt
	var err error

	// Make a statement!
	if stmt, err = f.db.Prepare("INSERT INTO Waypoint (Name, LocationId, Description, Difficulty) VALUES (?, ?, ?, ?)"); err != nil {
		log.Printf("Unable to prepare statement for creating waypoint %v", err)
		return err
	}
	defer stmt.Close()

	var result sql.Result
	if result, err = stmt.Exec(point.Name, point.Location.ID, point.Description, point.Difficulty); err != nil {
		log.Printf("Unable to create the startsite %v", err)
		return err
	}

	var siteid int64
	if siteid, err = result.LastInsertId(); err != nil {
		log.Printf("Cannot get the ID of the startsite %v", err)
		return err
	}

	point.ID = uint(siteid)

	return nil
}

// UpdateWayPoint - updates a waypoint might return FK err
func (f *MySQLRepository) UpdateWayPoint(ID uint, point *common.Waypoint) error {
	var stmt *sql.Stmt
	var err error

	// Make a statement!
	if stmt, err = f.db.Prepare("UPDATE Waypoint SET Name = ?, LocationId = ?, Description = ?, Difficulty = ? WHERE Id = ?"); err != nil {
		log.Printf("Unable to prepare update-waypoint statement %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(point.Name, point.Location.ID, point.Description, point.Difficulty, point.ID)

	return err
}

// GetWaypoint uses database views to get a singular waypoint
func (f *MySQLRepository) GetWaypoint(ID uint, point *common.Waypoint) error {
	var err error
	var stmt *sql.Stmt
	if stmt, err = f.db.Prepare("SELECT * FROM getwaypoint WHERE WaypointId = ?"); err != nil {
		log.Printf("Unable to prepare get-waypoint statement")
		return err
	}
	defer stmt.Close()

	return mapWaypoint(stmt.QueryRow(ID), point)
}

func mapWaypoint(row rowScanner, waypoint *common.Waypoint) error {
	if waypoint == nil {
		waypoint = &common.Waypoint{}
		waypoint.Location = common.Location{}
	}

	return row.Scan(&waypoint.Location.ID, &waypoint.Location.Elevation,
		&waypoint.Location.Description, &waypoint.Location.Name, &waypoint.Location.AreaName,
		&waypoint.Location.PostalCode, &waypoint.Location.CountryName,
		&waypoint.Location.Longitude, &waypoint.Location.Lattitude,
		&waypoint.ID, &waypoint.Name, &waypoint.Description, &waypoint.Difficulty)
}

// DeleteWaypoint soft-deletes a startsite
func (f *MySQLRepository) DeleteWaypoint(ID uint) error {
	var stmt *sql.Stmt
	var err error

	// Make a statement!
	if stmt, err = f.db.Prepare("UPDATE Waypoint SET Deleted = CURRENT_TIMESTAMP WHERE Id = ?"); err != nil {
		log.Printf("Unable to prepare delete start-site query %v", err)
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(ID)

	return err
}

// GetAllWaypoints get all waypoints in a paged and sivilized matter
func (f *MySQLRepository) GetAllWaypoints(limit int, page int) ([]common.Waypoint, error) {
	var stmt *sql.Stmt
	var err error

	if stmt, err = f.db.Prepare("SELECT * FROM getwaypoint LIMIT ?,?"); err != nil {
		log.Printf("Unable to prepare statement for retreiving all startsites %v", err)
		return nil, err
	}

	rows, err := stmt.Query(page*limit, limit)
	defer stmt.Close()

	if err != nil {
		log.Printf("Could not get any results %v", err)
		return nil, err
	}

	defer rows.Close()

	sites := make([]common.Waypoint, 0)

	for rows.Next() {
		var site common.Waypoint

		mapWaypoint(rows, &site)

		sites = append(sites, site)
	}
	return sites, nil
}

// GetStartSiteWaypoints gets all the waypoints for a given site
func (f *MySQLRepository) GetStartSiteWaypoints(siteID uint) ([]common.Waypoint, error) {
	var stmt *sql.Stmt
	var err error

	if stmt, err = f.db.Prepare("SELECT * FROM getwaypoint INNER JOIN StartsiteWaypoint ON getwaypoint.WaypointId = StartsiteWaypoint.waypointId WHERE StartsiteWaypoint.StartsiteId = ?;"); err != nil {
		log.Printf("Unable to prepare statement for retreiving all waypoints for one startsite %v", err)
		return nil, err
	}

	rows, err := stmt.Query(siteID)
	defer stmt.Close()

	if err != nil {
		log.Printf("Could not get any results %v", err)
		return nil, err
	}

	defer rows.Close()

	sites := make([]common.Waypoint, 0)

	for rows.Next() {
		var site common.Waypoint

		mapWaypoint(rows, &site)

		sites = append(sites, site)
	}
	return sites, nil
}
