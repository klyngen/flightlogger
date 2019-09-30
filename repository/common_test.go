package repository_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/klyngen/flightlogger/common"
	"github.com/klyngen/flightlogger/repository"
	"gotest.tools/assert"
)

func createConnection(t *testing.T) common.FlightLogDatabase {
	database := repository.MySQLRepository{}
	err := database.CreateConnection("root", "", "Flightlog", "", "")

	assert.NilError(t, err, "Could not connect to the database")

	return &database
}
