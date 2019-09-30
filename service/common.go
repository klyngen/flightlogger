package service

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"

	"github.com/klyngen/flightlogger/common"
	"github.com/klyngen/flightlogger/configuration"
)

// FlightLogService describes service containing business-logic
type FlightLogService struct {
	database   common.FlightLogDatabase
	config     configuration.ApplicationConfig
	email      common.EmailServiceInterface
	signingkey *rsa.PrivateKey
	verifykey  *rsa.PublicKey
}

// NewService creates a new flightlogservice
func NewService(database common.FlightLogDatabase, email common.EmailServiceInterface, config configuration.ApplicationConfig) *FlightLogService {
	sign, verify := getSigningKeys(config.PrivateKeyPath, config.PublicKeyPath)
	return &FlightLogService{database: database, config: config, signingkey: sign, verifykey: verify, email: email}
}

func getSigningKeys(privatekeypath string, publickeypath string) (*rsa.PrivateKey, *rsa.PublicKey) {
	var signBytes, verifyBytes []byte
	var signKey *rsa.PrivateKey
	var verifyKey *rsa.PublicKey
	var err error

	if signBytes, err = ioutil.ReadFile(privatekeypath); err != nil {
		log.Fatalf("Unable to read PrivateKey: %v", err)
	}

	if signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes); err != nil {
		log.Fatalf("Unable to parse PrivateKey: %v", err)
	}

	if verifyBytes, err = ioutil.ReadFile(publickeypath); err != nil {
		log.Fatalf("Unable to read PublicKey: %v", err)
	}

	if verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes); err != nil {
		log.Fatalf("Unable to parse PublicKey: %v", err)
	}
	return signKey, verifyKey
}

// VerifyTokenString - tries to verify the token using the certificate
func (f *FlightLogService) VerifyTokenString(tokenString string) (jwt.Claims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return f.verifykey, nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims, err
}
