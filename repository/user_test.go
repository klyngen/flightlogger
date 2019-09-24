package repository_test

import (
	"bytes"
	"log"
	"testing"

	"gotest.tools/assert"

	"github.com/klyngen/flightlogger/common"
)

func TestUserLifecycle(t *testing.T) {
	database := createConnection(t)

	user := common.User{
		FirstName:    "Ola",
		LastName:     "Nordmann",
		Email:        "ola@nordmann.no",
		PasswordHash: []byte("hashyHash"),
		PasswordSalt: []byte("reallySalty"),
	}

	assert.NilError(t, database.CreateUser(&user), "Could not create the user")

	// Assert that the data we intend to put in comes back the same way
	assert.Equal(t, "Ola", user.FirstName)
	assert.Equal(t, "Nordmann", user.LastName)
	assert.Equal(t, "ola@nordmann.no", user.Email)
	assert.Assert(t, (bytes.Compare(user.PasswordHash, []byte("hashyHash"))) == 0)
	assert.Assert(t, (bytes.Compare(user.PasswordSalt, []byte("reallySalty"))) == 0)

	// Change some data
	user.FirstName = "Kari"
	assert.NilError(t, database.UpdateUser(user.ID, &user), "Could not update a user")

	assert.Equal(t, "Kari", user.FirstName)

	// This user == nil. Should work anyways
	var newUser common.User
	assert.NilError(t, database.GetUser(user.ID, &newUser), "Could not get a single user")

	users, err := database.GetAllUsers(100, 0)

	assert.NilError(t, err, "Error returned when fetching multiple users")

	// We should have one or more users
	assert.Assert(t, len(users) > 0)

	var userByEmail common.User

	assert.NilError(t, database.GetUserByEmail(user.Email, &userByEmail), "Unable to fetch a user by email")

	var userByEmail2 common.User

	log.Println(database.GetUserByEmail("non existant", &userByEmail2))

	// TODO - implement testing for deletion of Many2Many tables

	// Cleanup after the test
	assert.NilError(t, database.DeleteUser(user.ID), "could not delete a user")
}
