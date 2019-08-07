package storage

import (
	"log"
	"testing"

	"github.com/klyngen/flightlogger/common"
)

func SetupDbTest() common.FlightLogDatabase {
	var db = &OrmDatabase{}

	err := db.CreateConnection("root", "Passwd", "flightlog", "3306", "localhost")

	if err != nil {
		log.Println(err)
	}

	return db
}

func TestMigration(t *testing.T) {
	db := SetupDbTest()

	err := db.MigrateDatabase()

	if err != nil {
		t.Fail()
	}
}
