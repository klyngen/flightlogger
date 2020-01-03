package service

import (
	"fmt"
	"log"
	"net/http"
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

type VerificationClaims struct {
	jwt.StandardClaims
	UserID string
}

// Authenticate should verify the logon and create a session
func (s *FlightLogService) Authenticate(username string, password string, request *http.Request) error {

	if err := s.sessionstore.RenewToken(request.Context()); err != nil {
		return errors.Wrap(err, "Unknown session error")
	}

	var user common.User
	// Try to get the user. It might not exist and might not be activated
	if err := s.database.GetUserByEmail(username, &user); err != nil {
		return errors.Wrap(err, "Could not fetch the user")
	}

	// Check if the hash is set
	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)) {
		return errors.Wrap(err, "Bad password")
	}


	// Try to set some session parameters
	s.sessionstore.Put(request.Context(), sessionParamUserID, user.ID)
	//s.sessionstore.Put(request.Context(), sessionParamRoles, user.)

	return nil
}

// ActivateUser is a straight pass-through to the database
func (s *FlightLogService) ActivateUser(UserID string) error {
	return s.database.ActivateUser(UserID)
}

// CreateVerificationToken creates a token used in an verification-email
func (s *FlightLogService) createVerificationToken(email string, ID string) (string, error) {
	expiration := time.Now().Add(time.Second * time.Duration(s.config.Tokenexpiration)).Unix()

	claims := &VerificationClaims{
		StandardClaims: jwt.StandardClaims{ExpiresAt: expiration},
		UserID:         ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(s.signingkey)
}

// CreateUser does not have any specific logic and will return the database-result / error
func (s *FlightLogService) CreateUser(user *common.User, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.PasswordHash = hash

	// Will create an inactive user
	if err = s.database.CreateUser(user); err != nil {
		log.Printf("we could not create the user: %s", err)
		return err
	}

	var tokenString string

	if tokenString, err = s.createVerificationToken(user.Email, user.ID); err != nil {
		log.Printf("Unable to create tokenstrings for verification for UID: %s. With error %v", user.ID, err)
		return err
	}

	log.Printf("User created with ID: %s,sending verification Email", user.ID)
	log.Println(s.email)
	// Try to send an verification-email to the user
	return s.email.SendVerificationEmail(user.Email,
		fmt.Sprintf("%s:%s/api/public/verify?token=%s", s.config.ServerURL, s.config.Serverport, tokenString))

}

// GetAllUsers does not have any specific logic and will return the database-result / error
func (s *FlightLogService) GetAllUsers(limit int, page int) ([]common.User, error) {
	return s.database.GetAllUsers(limit, page)
}

// GetUser does not have any specific logic and will return the database-result / error
func (s *FlightLogService) GetUser(ID string, user *common.User) error {
	return s.database.GetUser(ID, user)
}
