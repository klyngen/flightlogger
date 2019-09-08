package presentation

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klyngen/flightlogger/common"
)

type FlightLogApi struct {
	service common.FlightLogService
	router  *mux.Router
	port    string
}

func NewService(service common.FlightLogService, port string) FlightLogApi {
	router := mux.NewRouter()

	// Create the API
	return FlightLogApi{service: service, router: router, port: port}
}

func (api *FlightLogApi) StartApi() {
	err := http.ListenAndServe(fmt.Sprintf(":%s", api.port), api.router)

	if err != nil {
		log.Fatalf("Unable to start the API due to the following error: \n %v", err)
	}

	log.Printf("Started FlightLogger on port: %s", api.port)
}
