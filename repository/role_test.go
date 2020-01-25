package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/klyngen/flightlogger/common"
)

func TestRoleCreationCycle(t *testing.T) {
	db := createConnection(t)

	role := common.Role{
		Name:        "shitty role",
		Description: "some role used for bad jobs",
	}

	assert.NoError(t, db.CreateRole(&role), "Unable to create role")

	// ID should be set
	assert.NotEqual(t, 0, role.ID)

	assert.NoError(t, db.DeleteRole(role.ID))
}
