package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/klyngen/flightlogger/common"
)

// CreateWing -> remember to clean details before using the create
func (f *MySQLRepository) CreateWing(wing *common.FlyingDevice) error {
	stmt, err := f.db.Prepare("INSERT INTO Flightlog.FlyingDevice (Make, Model, FlyingDeviceTypeId) VALUES(?, ?, ?)")

	if err != nil {
		BadSqlError.NewFromException(err, "INSERT", "FlyingDevice")
	}

	// INSERT THE FLYING DEVICE
	res, err := stmt.Exec(wing.Make, wing.Model, wing.DeviceType)

	if err != nil {
		RowInsertionError.NewFromException(err, "INSERT", "FlyingDevice")
	}

	id, err := res.LastInsertId()

	if err != nil {
		RowInsertionError.New("Unable to get ID from rowInsertion", "INSERT", "FlyingDevice")
	}

	wing.ID = uint(id)

	return f.createDeviceDetails(wing.ID, wing.Details)
}

func (f *MySQLRepository) UpdateWing(ID uint, wing *common.FlyingDevice) error {
	stmt, err := f.db.Prepare("UPDATE Flightlog.FlyingDevice SET Make=?, Model=? WHERE Id=?")

	if err != nil {
		return BadSqlError.NewFromException(err, "UPDATE", "FlyingDevice")
	}

	_, err = stmt.Exec(wing.Make, wing.Model, wing.ID)

	if err != nil {
		return StatementExecutionError.NewFromException(err, "UPDATE", "FlyingDevice")
	}

	return nil
}

func (f *MySQLRepository) DeleteWing(ID uint) error {
	stmt, err := f.db.Prepare("DELETE FROM FlyingDevice WHERE Id = ? LIMIT 1")

	if err != nil {
		return BadSqlError.NewFromException(err, "DELETE", "FlyingDevice")
	}

	_, err = stmt.Exec(ID)

	if err != nil {
		return StatementExecutionError.NewFromException(err, "DELETE", "FlyingDevice")
	}

	return nil
}

func (f *MySQLRepository) GetWing(ID uint, wing *common.FlyingDevice) error {
	stmt, err := f.db.Prepare("SELECT DeviceId, DeviceMake, DeviceModel, DeviceTypeId FROM Flightlog.getflyingdevice WHERE DeviceID = ?; ")

	if err != nil {
		BadSqlError.NewFromException(err, "SELECT", "FlyingDevice")
	}

	row := stmt.QueryRow(ID)

	// Get the base data
	err = mapFlyingDevice(wing, row)

	if err != nil {
		return err // Already formatted
	}

	// Fetch the device details
	err = f.getDeviceDetails(wing)

	return err
}

func (f *MySQLRepository) GetAllWings(limit uint, page uint) ([]common.FlyingDevice, error) {
	sql := "SELECT DeviceId, DeviceMake, DeviceModel, DeviceTypeId FROM Flightlog.getflyingdevice LIMIT ?,?; "
	stmt, err := f.db.Prepare(sql)

	if err != nil {
		return nil, BadSqlError.NewFromException(err, "SELECT", "FlyingDevice")
	}

	rows, err := stmt.Query((page-1)*limit, limit)

	if err != nil {
		return nil, TransactionError.NewFromException(err, "SELECT", "FlyingDevice")
	}

	devices := make([]common.FlyingDevice, 0)
	for rows.Next() {
		var device common.FlyingDevice
		if err := mapFlyingDevice(&device, rows); err != nil {
			log.Printf("Unable to map data from request '%s' with parameters %d %d", sql, (page-1)*limit, limit)
		} else {
			devices = append(devices, device)
		}
	}

	return devices, err
}

func (f *MySQLRepository) GetWingSearchByName(name string, wings []common.FlyingDevice) error {
	panic("not implemented")
}

// Creates details for a flyingDevice
func (f *MySQLRepository) createDeviceDetails(ID uint, details []common.FlyingDeviceDetails) error {

	tx, err := f.db.BeginTx(context.Background(), &sql.TxOptions{})

	if err != nil {
		return DriverFunctionError.NewFromException(err, "INSERT", "FlyingDevice")
	}

	stmt, err := tx.Prepare("INSERT INTO Flightlog.FlyingDeviceDetails (FlyingDeviceId, DetailName, DetailDescription) VALUES(?, ?, ?)")

	if err != nil {
		return BadSqlError.NewFromException(err, "INSERT", "FlyingDevice")
	}

	defer stmt.Close()

	for _, d := range details {
		stmt.Exec(ID, d.DetailName, d.Description)
	}

	err = tx.Commit()

	if err != nil {
		TransactionError.NewFromException(err, "INSERT", "FlyingDevice")
	}

	return nil
}

func (f *MySQLRepository) getDeviceDetails(wing *common.FlyingDevice) error {
	stmt, err := f.db.Prepare("SELECT Id, DetailName, DetailDescription FROM Flightlog.FlyingDeviceDetails WHERE FlyingDeviceId = ? AND DELETED IS NULL")

	if err != nil {
		return BadSqlError.NewFromException(err, "SELECT", "FlyingDeviceDetails")
	}

	rows, err := stmt.Query(wing.ID)

	for rows.Next() {
		var details common.FlyingDeviceDetails

		err = mapFlyingDeviceDetails(&details, rows)

		if err != nil {
			log.Printf("Unable to map DeviceDetails for flyingDevice %d and device details %d", wing.ID, details.ID)
			err = nil
		} else { // Everything is OK
			wing.Details = append(wing.Details, details)
		}
	}

	return nil

}

func mapFlyingDevice(device *common.FlyingDevice, row rowScanner) error {
	if device == nil {
		device = &common.FlyingDevice{
			Details: make([]common.FlyingDeviceDetails, 0),
		}
	}
	err := row.Scan(&device.ID, &device.Make, &device.Model, &device.DeviceType)

	if err != nil {
		SerilizationError.NewFromException(err, "SELECT", "FlyingDevice")
	}

	return nil
}

func mapFlyingDeviceDetails(details *common.FlyingDeviceDetails, scanner rowScanner) error {
	if details == nil {
		details = &common.FlyingDeviceDetails{}
	}

	err := scanner.Scan(&details.ID, &details.DetailName, &details.Description)

	if err != nil {
		return SerilizationError.NewFromException(err, "SELECT", "FlyingDeviceDetail")
	}

	return nil
}
