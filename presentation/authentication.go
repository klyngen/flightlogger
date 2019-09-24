package presentation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/klyngen/flightlogger/common"
	"github.com/klyngen/jsend"

	"github.com/dgrijalva/jwt-go"
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

	token, err := f.service.Authenticate(creds.Username, creds.Password)

	if err != nil {
		jsend.FormatResponse(w, "Bad credentials", jsend.UnAuthorized)
		return
	}
	expiration := time.Now().Add(time.Second * time.Duration(f.timeout))

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expiration,
	})

	jsend.FormatResponse(w, "Success", jsend.Success)
}

// Middleware to verify the accesstoken
func (f *FlightLogApi) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			jsend.FormatResponse(w, "You have no accesstoken. RTFM", jsend.UnAuthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		_, err := verifyToken(tokenString, f.secret)
		if err != nil {
			jsend.FormatResponse(w, "Error verifying token. Expired or invalid", jsend.UnAuthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getToken(user common.User) (string, error) {
	signingKey := []byte("keymaker")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		"email": user.Email,
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func verifyToken(tokenString string, secret string) (jwt.Claims, error) {
	signingKey := []byte(secret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims, err
}
