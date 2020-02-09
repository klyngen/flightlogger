package main

import (
	"fmt"
	"log"

	"github.com/alexedwards/scs/redisstore"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/klyngen/flightlogger/common"
	"github.com/klyngen/flightlogger/email"
	"github.com/klyngen/flightlogger/service"

	"github.com/klyngen/flightlogger/repository"

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

	// Create the database connection (DataLayer)
	err := db.CreateConnection(config.DatabaseConfiguration.Username,
		config.DatabaseConfiguration.Password,
		config.DatabaseConfiguration.Database,
		config.DatabaseConfiguration.Port,
		config.DatabaseConfiguration.Hostname)

	if err != nil {
		log.Fatalf("Likely a database misconfiguration: %v", err)
	}

	// Should be enough to add email-support to our application (DataLayer)
	emailService := email.NewEmailService(config.EmailConfiguration)

	if emailService == nil {
		panic("Cannot have non-existing email-service")
	}

	// Create the casbin-adapter
	adapter, _ := xormadapter.NewAdapter("mysql",
		createConnectionString(config.DatabaseConfiguration.Username,
			config.DatabaseConfiguration.Password,
			config.DatabaseConfiguration.Hostname,
			config.DatabaseConfiguration.Port))

	var fservice common.FlightLogService

	// Instantiate our use-case / service-layer
	if config.RedisConfiguration.IsEmpty() {
		fservice = service.NewService(db, emailService, config, adapter)
	} else {
		redisPool := createRedisPool(config.RedisConfiguration)
		fservice = service.NewServiceWithPersistedSession(db, emailService, config, redisstore.New(redisPool), adapter)
	}

	// Create our presentation layer
	api := presentation.NewService(fservice, config)
	api.StartAPI() // LET THE GAMES BEGIN
}

func createRedisPool(config configuration.DatabaseConfig) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%s", config.Hostname, config.Port),
				redis.DialPassword(config.Password),
			)
		},
	}
}

func createConnectionString(username string, password string, hostname string, port string) string {
	log.Println(port)
	if len(hostname) > 0 { // Full config
		return fmt.Sprintf("%v:%v@tcp(%v:%v)/", username, password, hostname, port)
	}

	return fmt.Sprintf("%v:%v@/", username, password)
}
