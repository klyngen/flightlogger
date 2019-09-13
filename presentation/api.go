package presentation

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klyngen/flightlogger/common"
	"github.com/klyngen/jsend"
)

type FlightLogApi struct {
	service common.FlightLogService
	router  *mux.Router
	port    string
}

func NewService(service common.FlightLogService, port string) FlightLogApi {
	router := mux.NewRouter().PathPrefix("/api/protected").Subrouter()
	unprotected := router.PathPrefix("/api/public").Subrouter()
	// Mount authenticationRoutes
	mountAuthenticationRoutes(unprotected)

	// Jsendify the default handlers
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(notAllowedHandler)

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

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	jsend.FormatResponse(w, "No such endpoint. RTFM", jsend.NotFound)
}

func notAllowedHandler(w http.ResponseWriter, r *http.Request) {
	jsend.FormatResponse(w, "Correct endpoint wrong method. RTFM", jsend.MethodNotAllowed)
}
