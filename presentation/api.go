package presentation

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/klyngen/flightlogger/configuration"

	"github.com/gorilla/mux"
	"github.com/klyngen/flightlogger/common"
	"github.com/klyngen/jsend"
)

// FlightLogApi describes a presentation-service
type FlightLogApi struct {
	service common.FlightLogService
	router  *mux.Router
	secret  string
	timeout int
	port    string
}

// NewService creates an api with defined routes
func NewService(service common.FlightLogService, config configuration.ApplicationConfig) FlightLogApi {
	router := mux.NewRouter()

	unprotected := router.PathPrefix("/api/public").Subrouter()
	protected := router.PathPrefix("/api/protected").Subrouter()
	// Mount authenticationRoutes

	// Jsendify the default handlers
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(notAllowedHandler)

	// Create the API
	api := FlightLogApi{service: service, router: router, port: config.Serverport}
	api.mountAuthenticationRoutes(unprotected)
	api.mountUserRoutes(protected)

	// Middleware to require login for certain endpoints
	protected.Use(api.authMiddleware)

	// TODO: add authorization

	return api
}

// StartAPI does just that and prints routes
func (api *FlightLogApi) StartAPI() {
	printRoutes(api.router)

	log.Printf("Started FlightLogger on port: %s", api.port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", api.port), api.router)

	if err != nil {
		log.Fatalf("Unable to start the API due to the following error: \n %v", err)
	}

}

func printRoutes(router *mux.Router) {
	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			log.Println("ROUTE:", pathTemplate)
		}
		methods, err := route.GetMethods()
		if err == nil {
			log.Println("Methods:", strings.Join(methods, ","))
		}

		log.Println()
		return nil
	})

	if err != nil {
		log.Println(err)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	jsend.FormatResponse(w, "No such endpoint. RTFM", jsend.NotFound)
}

func notAllowedHandler(w http.ResponseWriter, r *http.Request) {
	jsend.FormatResponse(w, "Correct endpoint wrong method. RTFM", jsend.MethodNotAllowed)
}
