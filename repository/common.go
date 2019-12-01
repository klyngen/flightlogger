package repository

import (
	"database/sql"
	"fmt"
)

// MySQLRepository describes how we interact with the database
type MySQLRepository struct {
	db *sql.DB
}

// CreateConnection What all databases should do
func (f *MySQLRepository) CreateConnection(username string, password string, database string, port string, hostname string) error {

	var db *sql.DB
	var err error
	if len(hostname) > 0 { // Full config
		db, err = sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", username, password, hostname, port, database))
	} else { // Simple config
		db, err = sql.Open("mysql", fmt.Sprintf("%v:%v@/%v", username, password, database))
	}

	// Format the string and connect to the database
	// Set the databaseobject
	f.db = db
	return err
}

type rowScanner interface {
	Scan(args ...interface{}) error
}

// DEFINE COMMON ERRORS
type DataLayerError struct {
	message   string // Example -> foreignkeyissue
	action    string // Example -> INSERT
	entity    string // Example -> FlyingDevice
	errorType DataLayerErrorType
}

func (e *DataLayerError) Error() string {
	return fmt.Sprintf("'%s' to '%s' created the following error-message: '%s'", e.action, e.entity, e.message)
}

func (e *DataLayerError) Type() DataLayerErrorType {
	return e.errorType
}

type DataLayerErrorType int

// New creates a new error
func (t DataLayerErrorType) New(message string, action string, entity string) *DataLayerError {
	return &DataLayerError{
		message:   message,
		action:    action,
		entity:    entity,
		errorType: t,
	}
}

// NewFromException - not neccessary but makes code simpler
func (t DataLayerErrorType) NewFromException(err error, action string, entity string) *DataLayerError {
	return t.New(err.Error(), action, entity)
}

const (
	BadSqlError             DataLayerErrorType = 1
	EntityResolutionError   DataLayerErrorType = 2
	RowInsertionError       DataLayerErrorType = 3
	SerilizationError       DataLayerErrorType = 4
	DriverFunctionError     DataLayerErrorType = 5
	TransactionError        DataLayerErrorType = 6
	StatementExecutionError DataLayerErrorType = 7
)
