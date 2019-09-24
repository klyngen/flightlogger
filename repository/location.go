package repository

import "github.com/klyngen/flightlogger/common"

// Location CRUD and search
func (f *MySQLRepository) CreateLocation(location *common.Location) error {
	panic("not implemented")
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
