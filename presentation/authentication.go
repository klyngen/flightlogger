package presentation

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/klyngen/flightlogger/common"
	"github.com/klyngen/jsend"

	"github.com/gorilla/mux"
)

type credentials struct {
	Username string
	Password string
}

type registerContent struct {
	common.User
	Password string
}

func (f *FlightLogApi) mountAuthenticationRoutes(router *mux.Router) {
	router.HandleFunc("/login", f.loginHandler).Methods("POST")
	router.HandleFunc("/createuser", f.newUserHandler).Methods("POST")
	router.HandleFunc("/verify", f.verifyUserAccount).Methods("GET")
}

// TODO: make this redirect to some GUI
func (f *FlightLogApi) verifyUserAccount(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query()["token"]
	log.Printf("Get user for ID: %v", token)

	if len(token[0]) > 0 {
		claims, err := f.service.VerifyTokenString(token[0])

		if err != nil {
			log.Printf("Invalid token passed to service %v", token[0])
			jsend.FormatResponse(w, "Bad token", jsend.BadRequest)
			return
		}

		// "Parse" the claims
		userID := claims.(jwt.MapClaims)["UserID"]
		log.Println(claims, userID)

		if err = f.service.ActivateUser(userID.(string)); err != nil {
			log.Printf("Unable to activate userID %s, due to erro %v", userID, err)
			jsend.FormatResponse(w, "Could not activate the user", jsend.InternalServerError)
			return
		}

		jsend.FormatResponse(w, "User is activated", jsend.Success)

		return
	}

	jsend.FormatResponse(w, "No token is present. Are you trying to hack me!?", jsend.BadRequest)

}

func (f *FlightLogApi) newUserHandler(w http.ResponseWriter, r *http.Request) {
	var user registerContent

	// If we cannot decode the request
	if json.NewDecoder(r.Body).Decode(&user) != nil {
		jsend.FormatResponse(w, "Bad request data. RTFM", jsend.BadRequest)
		return
	}

	// We should be able to pass the error without it being too dirty
	err := f.service.CreateUser(&user.User, user.Password)

	if err != nil {
		jsend.FormatResponse(w, err.Error(), jsend.BadRequest)
		return
	}

	jsend.FormatResponse(w, "User is created", jsend.NoContent)

}

func (f *FlightLogApi) loginHandler(w http.ResponseWriter, r *http.Request) {
	var creds credentials

	// If we cannot decode the request
	if json.NewDecoder(r.Body).Decode(&creds) != nil {
		jsend.FormatResponse(w, "Bad request data. RTFM", jsend.BadRequest)
		return
	}

	err := f.service.Authenticate(creds.Username, creds.Password, r)

	if err != nil {
		jsend.FormatResponse(w, err.Error(), jsend.UnAuthorized)
		return
	}

	jsend.FormatResponse(w, "Authenticated!", jsend.Success)
}

// Middleware to verify the accesstoken
func (f *FlightLogApi) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Verify the session

		session := f.service.GetSessionManager()

		userID := session.GetString(r.Context(), string(common.SessionParamUserID))

		// If the session has a userID we let them pass
		if len(userID) == 0 {
			// Give the middle finger if there is no userID
			//jsend.FormatResponse(w, "Unauthorized", jsend.UnAuthorized)
			http.Error(w, "Missing cookie", http.StatusUnauthorized)
		} else {
			role := session.GetString(r.Context(), string(common.SessionParamRoles))

			if role == "" {
				role = "anonymous"
			}

			enforcer := f.service.GetCasbinEnforcer()

			res, err := enforcer.Enforce(role, userID, r.RequestURI, r.Method)

			if err != nil {
				//jsend.FormatResponse(w, "Authorization error", jsend.InternalServerError)
				log.Println(err)
				http.Error(w, "Internal error", http.StatusInternalServerError)
			} else if res {
				next.ServeHTTP(w, r) // YOU SHALL PASS
				return
			}
		}

		jsend.FormatResponse(w, "You cannot access this endpoint", jsend.Forbidden)
		//http.Error(w, "Forbidden", http.StatusForbidden)
	})
}

// Middleware to verify the accesstoken
func (f *FlightLogApi) jwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			jsend.FormatResponse(w, "You have no accesstoken. RTFM", jsend.UnAuthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		_, err := f.service.VerifyTokenString(tokenString)
		if err != nil {
			jsend.FormatResponse(w, "Error verifying token. Expired or invalid", jsend.UnAuthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
