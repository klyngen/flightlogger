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
