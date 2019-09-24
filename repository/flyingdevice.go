package repository

import "github.com/klyngen/flightlogger/common"

// Wing CRUD
func (f *MySQLRepository) CreateWing(wing *common.Wing) error {
	panic("not implemented")
}

func (f *MySQLRepository) UpdateWing(ID uint, wing *common.Wing) error {
	panic("not implemented")
}

func (f *MySQLRepository) DeleteWing(ID uint) error {
	panic("not implemented")
}

func (f *MySQLRepository) GetWing(ID uint, wing *common.Wing) (common.Wing, error) {
	panic("not implemented")
}

func (f *MySQLRepository) GetAllWings(limit uint, page uint, wing *common.Wing) error {
	panic("not implemented")
}

func (f *MySQLRepository) GetWingSearchByName(name string, wings []common.Wing) error {
	panic("not implemented")
}
