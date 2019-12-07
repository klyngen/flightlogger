package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/klyngen/flightlogger/common"
)

// CreateFlight creates a new flight
func (f *MySQLRepository) CreateFlight(flight *common.Flight) error {
	stmt, err := f.db.Prepare("INSERT INTO Flightlog.Flight (Id, UserId, StartsiteId, WaypointId, Duration, Distance, MaxHeight, Hangtime, FlightDescription, FlyingDeviceId) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return BadSqlError.NewFromException(err, "INSERT", "Flight")
	}
	defer stmt.Close()

	// This is now NULL
	var waypointID interface{}

	if flight.Waypoint.ID != 0 {
		waypointID = flight.Waypoint.ID
	}

	flight.ID = guidMaker()

	res, err := stmt.Exec(flight.ID, flight.User.ID, flight.Startsite.ID,
		waypointID, flight.Duration, flight.Distance,
		flight.MaxHight, flight.HangTime, flight.Notes, flight.Wing.ID)

	if err != nil {
		flight.ID = ""
		return StatementExecutionError.NewFromException(err, "INSERT", "Flight")
	}

	if affected, err := res.RowsAffected(); err == nil {
		if affected == 0 {
			return RowInsertionError.New("No rows are actually inserted", "INSERT", "FLIGHT")
		}
	}

	return nil
}

// UpdateFlight updates only the flight-entity
func (f *MySQLRepository) UpdateFlight(ID string, flight *common.Flight) error {
	//f.db.Prepare()
	stmt, err := f.db.Prepare("UPDATE Flightlog.Flight SET UserId=?, StartsiteId=?, WaypointId=?, Duration=?, Distance=?, MaxHeight=?, Hangtime=?, FlightDescription=?, FlyingDeviceId=? WHERE Id=?; ")

	if err != nil {
		return BadSqlError.NewFromException(err, "UPDATE", "Flight")
	}

	defer stmt.Close()

	var waypointID interface{}

	if flight.Waypoint.ID != 0 {
		waypointID = flight.Waypoint.ID
	}

	_, err = stmt.Exec(flight.User.ID, flight.Startsite.ID, waypointID, flight.Duration, flight.Distance, flight.MaxHight, flight.HangTime, flight.Notes, flight.Wing.ID, flight.ID)

	if err != nil {
		return StatementExecutionError.NewFromException(err, "UPDATE", "Flight")
	}

	return nil
}

// DeleteFlight deletes a flight either soft or hard. Depends on user wish
func (f *MySQLRepository) DeleteFlight(ID string, soft bool) error {
	var stmt *sql.Stmt
	var err error

	if soft { // Hope its not microsoft... Then we are in trouble
		stmt, err = f.db.Prepare("DELETE FROM Flight WHERE Id = ? LIMIT 1")
	} else { // Do it hard ;)
		stmt, err = f.db.Prepare("UPDATE Flight SET Deleted = current_timestamp() WHERE Id = ? LIMIT 1")
	}

	if err != nil {
		BadSqlError.NewFromException(err, "DELETE", "Flight")
	}
	defer stmt.Close()

	_, err = stmt.Exec(ID)

	if err != nil {
		return StatementExecutionError.NewFromException(err, "DELETE", "Flight")
	}

	return nil
}

func (f *MySQLRepository) GetAllFlights(limit int, page int) ([]common.Flight, error) {
	stmt, err := f.db.Prepare("SELECT * FROM getflight LIMIT ?,?")

	if err != nil {
		return nil, BadSqlError.NewFromException(err, "SELECT", "Flight")
	}

	res, err := stmt.Query((page-1)*limit, limit)

	if err != nil {
		return nil, StatementExecutionError.NewFromException(err, "SELECT", "Flight")
	}

	results := make([]common.Flight, 0)

	for res.Next() {
		var flight common.Flight

		if err = mapFlight(&flight, res); err != nil {
			log.Printf("Unable to serialize Flight %v", err)
		} else {
			results = append(results, flight)
		}
	}

	return results, err
}

// GetFlight gets a flight based on ID using a view
func (f *MySQLRepository) GetFlight(ID string, flight *common.Flight) error {

	stmt, err := f.db.Prepare("SELECT * from getflight WHERE Id = ? LIMIT 1")

	if err != nil {
		return BadSqlError.NewFromException(err, "SELECT", "Flight")
	}

	defer stmt.Close()

	row := stmt.QueryRow(ID)

	return mapFlight(flight, row)
}

func mapFlight(flight *common.Flight, scanner rowScanner) error {
	if flight == nil {
		flight = &common.Flight{
			User:      common.User{},
			Wing:      common.FlyingDevice{},
			Startsite: common.StartSite{},
			Waypoint:  common.Waypoint{},
		}
	}

	var waypointID interface{}
	var created string

	if err := scanner.Scan(&flight.ID, &flight.Wing.ID, &flight.User.ID, &flight.Startsite.ID,
		&waypointID, &flight.Distance, &flight.Duration, &flight.MaxHight,
		&flight.HangTime, &created); err != nil {

		return SerilizationError.NewFromException(err, "SELECT", "Flight")
	}

	if waypointID != nil {
		flight.Waypoint.ID = waypointID.(uint)
	}

	flight.Created, _ = time.Parse("2006-02-01 15:04:05", created)

	return nil
}

// FlightIncident CRUD and search
func (f *MySQLRepository) CreateFlightIncident(incident *common.Incident) error {
	panic("not implemented")
}

func (f *MySQLRepository) UpdateFlightIncident(ID uint, incident *common.Incident) error {
	panic("not implemented")
}

func (f *MySQLRepository) DeleteFlightIncident(ID uint) error {
	panic("not implemented")
}

func (f *MySQLRepository) GetFlightIncident(ID uint, flight *common.Flight) error {
	panic("not implemented")
}

func (f *MySQLRepository) GetFlightIncidentByLevel(errorLevel uint, flights []common.Flight) error {
	panic("not implemented")
}

func (f *MySQLRepository) GetFlightIncidents(limit int, page int, flights []common.Flight) error {
	panic("not implemented")
}
