package main

import (
	"fmt"
	"log"
	"github.com/klyngen/flightlogger/configuration"
	"github.com/klyngen/flightlogger/presentation"
	"github.com/klyngen/flightlogger/storage"
	"github.com/klyngen/flightlogger/usecase"
)

func main() {
	fmt.Println("##### STARTING FLIGHTLOG BACKEND ####")
	log.Println("Starting flightlog API")
	log.

	// ######## BUILD THE SERVICE ##############

	// Get the configuration - WILL PANIC IF FAILS
	config := configuration.GetConfiguration()

	db := &storage.OrmDatabase{}
	
	db.CreateConnection(config.DatabaseConfiguration.Username,
		config.DatabaseConfiguration.Password,
		config.DatabaseConfiguration.Database,
		config.DatabaseConfiguration.Port,
		config.DatabaseConfiguration.Hostname)

	service := usecase.NewService(db)

	api := presentation.NewService(service, config.Serverport)
	api.StartApi()
}

