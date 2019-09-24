package repository

import "github.com/klyngen/flightlogger/common"

// FileCreation CRUD
func (f *MySQLRepository) CreateFile(file *common.FileReference) error {
	panic("not implemented")
}

func (f *MySQLRepository) GetFile(ID uint, file *common.FileReference) error {
	panic("not implemented")
}

func (f *MySQLRepository) DeleteFile(ID uint) error {
	panic("not implemented")
}
