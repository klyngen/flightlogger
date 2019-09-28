package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/klyngen/flightlogger/common"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Claims describes the basic claims in this plattform
type Claims struct {
	Username  string
	UserID    string
	Firstname string
	Lastname  string
	jwt.StandardClaims
}

// Authenticate should verify the logon and create a valid JWT-token
func (s *FlightLogService) Authenticate(username string, password string) (string, error) {
	var user common.User
	// Get the user
	err := s.database.GetUserByEmail(username, &user)

	if err != nil {

		return "", errors.Wrap(err, "Could not fetch the user")
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))

	if err != nil {
		return "", errors.Wrap(err, "Bad password")
	}

	expiration := time.Now().Add(time.Second * time.Duration(s.config.Tokenexpiration)).Unix()
	// Create the JWT-token
	claims := &Claims{
		Username:       user.Email,
		Firstname:      user.FirstName,
		Lastname:       user.LastName,
		UserID:         user.ID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expiration},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(s.signingkey)

	if err != nil {
		return "", errors.Wrap(err, "Unable to create token")
	}

	return tokenString, nil
}

// CreateUser does not have any specific logic and will return the database-result / error
func (s *FlightLogService) CreateUser(user *common.User, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.PasswordHash = hash
	user.PasswordSalt = []byte("sdfsdf")

	return s.database.CreateUser(user)
}

// GetAllUsers does not have any specific logic and will return the database-result / error
func (s *FlightLogService) GetAllUsers(limit int, page int) ([]common.User, error) {
	return s.database.GetAllUsers(limit, page)
}

// GetUser does not have any specific logic and will return the database-result / error
func (s *FlightLogService) GetUser(ID string, user *common.User) error {
	return s.database.GetUser(ID, user)
}
