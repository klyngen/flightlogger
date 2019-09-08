package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/klyngen/flightlogger/configuration"
	"github.com/klyngen/flightlogger/storage"
)

func main() {
	fmt.Println("##### STARTING FLIGHTLOG BACKEND ####")
	log.Println("Starting flightlog API")

	// ######## BUILD THE SERVICE ##############

	// Get the configuration - WILL PANIC IF FAILS
	config := configuration.GetConfiguration()

	db := &storage.OrmDatabase{}
	db.CreateConnection(config.DatabaseConfiguration.Username, config.DatabaseConfiguration.Password, config.DatabaseConfiguration.Database, config.DatabaseConfiguration.Port, config.DatabaseConfiguration.Hostname)

	// Listen to the configured port
	http.ListenAndServe(fmt.Sprintf(":%s", config.Serverport), application.Api)
}
