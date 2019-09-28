package presentation

import (
	"log"
	"net/http"

	"github.com/klyngen/jsend"

	"github.com/gorilla/mux"
	"github.com/klyngen/flightlogger/common"
)

func (f *FlightLogApi) mountUserRoutes(router *mux.Router) {
	router.HandleFunc("/user", f.getUser).Methods("GET")
}

// Get a user from the api
func (f *FlightLogApi) getUser(w http.ResponseWriter, r *http.Request) {
	userid := r.URL.Query()["uid"]

	log.Printf("Get user for ID: %v", userid[0])
	if len(userid[0]) > 0 {
		var user common.User

		// If it does not work
		if err := f.service.GetUser(userid[0], &user); err != nil {
			log.Printf("Unable to get userId: %s, got error %v", userid[0], err)
			jsend.FormatResponse(w, "Unable to fetch user", jsend.InternalServerError)
			return
		}

		// If it worked
		if &user == nil {
			log.Println("this was wrong")
			jsend.FormatResponse(w, nil, jsend.NoContent)
			return
		}

		jsend.FormatResponse(w, user, jsend.Success)
		return
	}

	// The userId is empty
	jsend.FormatResponse(w, "No userid given. uid-parameter must be set", jsend.BadRequest)
}
