package service

import (
	"github.com/klyngen/flightlogger/common"
	"github.com/klyngen/flightlogger/configuration"
)

// FlightLogService describes service containing business-logic
type FlightLogService struct {
	database common.FlightLogDatabase
	config   configuration.ApplicationConfig
}

// NewService creates a new flightlogservice
func NewService(database common.FlightLogDatabase, config configuration.ApplicationConfig) *FlightLogService {
	return &FlightLogService{database: database, config: config}
}
