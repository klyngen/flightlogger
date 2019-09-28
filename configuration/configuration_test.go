package configuration

import (
	"fmt"
	"os"
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

	assert.Equal(t, "42", config.EmailConfiguration.Port)
	assert.Equal(t, "uname", config.EmailConfiguration.Username)
	assert.Equal(t, "pass", config.EmailConfiguration.Password)
	assert.Equal(t, "server", config.EmailConfiguration.SMTPServer)
}

func TestGettingAbsolutePath(t *testing.T) {
	homedir, _ := os.UserHomeDir()
	testPathTilde := "~/somepath"
	testPathHome := "$HOME/somepath"

	absolutePath := "/someAbsolutePath/somewhere"
	relativePath := "../someAbsolutePath/somewhere"

	resultPath := fmt.Sprintf("%v/%v", homedir, "somepath")

	assert.Equal(t, resultPath, getAbsoulePath(testPathTilde))
	assert.Equal(t, resultPath, getAbsoulePath(testPathHome))
	assert.Equal(t, absolutePath, getAbsoulePath(absolutePath))
	assert.Assert(t, relativePath != getAbsoulePath(relativePath))

}
