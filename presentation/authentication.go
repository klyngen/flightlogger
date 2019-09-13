package presentation

import (
	"net/http"
	"strings"

	"github.com/klyngen/jsend"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func mountAuthenticationRoutes(router *mux.Router) {

}

// Middleware to verify claims sets the needed claims
func authMiddleware(next http.Handler, validClaims map[string]string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			jsend.FormatResponse(w, "You have no accesstoken. RTFM", jsend.UnAuthorized)
			return
		}
		
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := verifyToken(tokenString)
		if err != nil {
			jsend.FormatResponse(w, "Error verifying token. Expired or invalid", jsend.UnAuthorized)
			return
		}

		mappedClaims := claims.(map[string]string)

		if !authorize(validClaims, mappedClaims) {
			jsend.FormatResponse(w, "You have no access to that endpoint", jsend.Forbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func authorize(validClaims map[string]string, givenClaims map[string]string) bool {

	return true
}

func getToken(name string) (string, error) {
	signingKey := []byte("keymaker")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"role": "redpill",
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("keymaker")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims, err
}
