package main

import (
	"fmt"
	"log"

	"github.com/klyngen/flightlogger/repository"
	"github.com/klyngen/flightlogger/service"

	"github.com/klyngen/flightlogger/configuration"
	"github.com/klyngen/flightlogger/presentation"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("##### STARTING FLIGHTLOG BACKEND ####")
	log.Println("Starting flightlog API")

	// ######## BUILD THE SERVICE ##############

	// Get the configuration - WILL PANIC IF FAILS
	config := configuration.GetConfiguration()

	db := &repository.MySQLRepository{}

	err := db.CreateConnection(config.DatabaseConfiguration.Username,
		config.DatabaseConfiguration.Password,
		config.DatabaseConfiguration.Database,
		config.DatabaseConfiguration.Port,
		config.DatabaseConfiguration.Hostname)

	if err != nil {
		log.Fatalf("Likely a database misconfiguration: %v", err)
	}

	log.Println(db)

	service := service.NewService(db, config)
	api := presentation.NewService(service, config)
	api.StartAPI()
}
