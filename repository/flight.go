package repository

import "github.com/klyngen/flightlogger/common"

// Flight CRUD
func (f *MySQLRepository) CreateFlight(flight *common.Flight) error {
	panic("not implemented")
}

func (f *MySQLRepository) UpdateFlight(ID uint, flight *common.Flight) error {
	panic("not implemented")
}

func (f *MySQLRepository) DeleteFlight(ID uint, soft bool) error {
	panic("not implemented")
}

func (f *MySQLRepository) GetAllFlights(limit int, page int, flights []common.Flight) error {
	panic("not implemented")
}

func (f *MySQLRepository) GetFlight(ID uint, flight *common.Flight) error {
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
