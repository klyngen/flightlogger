package repository_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/klyngen/flightlogger/repository"
)

func TestErrorCreation(t *testing.T) {
	err := repository.BadSqlError.New("test exception", "INSERT", "FlyingDevice")

	assert.Equal(t, repository.BadSqlError, err.Type())
	assert.Equal(
		t,
		"'INSERT' to 'FlyingDevice' created the following error-message: 'test exception'",
		err.Error(),
	)
}
