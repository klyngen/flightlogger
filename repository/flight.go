package repository

import (
	"database/sql"

	"github.com/klyngen/flightlogger/common"
)

// CreateFlight creates a new flight
func (f *MySQLRepository) CreateFlight(flight *common.Flight) error {
	stmt, err := f.db.Prepare("INSERT INTO Flightlog.Flight (Id, UserId, StartsiteId, WaypointId, Duration, Distance, MaxHeight, Hangtime, FlightDescription, FlyingDeviceId) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return BadSqlError.NewFromException(err, "INSERT", "Flight")
	}
	stmt.Close()

	flight.ID = guidMaker()

	_, err = stmt.Exec(flight.ID, flight.User.ID, flight.Startsite.ID,
		flight.Waypoint.ID, flight.Duration, flight.Distance,
		flight.MaxHight, flight.HangTime, flight.Notes, flight.Wing.ID)

	if err != nil {
		StatementExecutionError.NewFromException(err, "INSERT", "Flight")
	}

	return nil
}

// UpdateFlight updates only the flight-entity
func (f *MySQLRepository) UpdateFlight(ID uint, flight *common.Flight) error {
	//f.db.Prepare()
	stmt, err := f.db.Prepare("UPDATE Flightlog.Flight SET UserId=?, StartsiteId=?, WaypointId=?, Duration=?, Distance=?, MaxHeight=?, Hangtime=?, FlightDescription=?, FlyingDeviceId=?, WHERE Id=?; ")

	if err != nil {
		return BadSqlError.NewFromException(err, "UPDATE", "Flight")
	}

	_, err = stmt.Exec(flight.User.ID, flight.Startsite.ID, flight.Waypoint.ID, flight.Duration, flight.Distance, flight.MaxHight, flight.HangTime, flight.ID)

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

	_, err = stmt.Exec(ID)

	if err != nil {
		return StatementExecutionError.NewFromException(err, "DELETE", "Flight")
	}

	return nil
}

func (f *MySQLRepository) GetAllFlights(limit int, page int, flights []common.Flight) error {
	panic("not implemented")
}

func (f *MySQLRepository) GetFlight(ID string, flight *common.Flight) error {
	panic("not implemented")
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
