package configuration

import (
	"testing"

	"gotest.tools/assert"
)

/*
	THIS TEST IS OVERKILL BUT IT IS TO
	ENSURE THAT THE TASK IS COMPLETED
*/

func TestConfigurationUnmarshalling(t *testing.T) {
	config := GetConfiguration()

	assert.Equal(t, "61225", config.Serverport)

	assert.Equal(t, "hostname", config.DatabaseConfiguration.Hostname)
	assert.Equal(t, "password", config.DatabaseConfiguration.Password)
	assert.Equal(t, "port", config.DatabaseConfiguration.Port)
	assert.Equal(t, "username", config.DatabaseConfiguration.Username)

}
