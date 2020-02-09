package service

import "testing"

import "github.com/stretchr/testify/assert"

func TestIsOwner(t *testing.T) {
	// Happy day case
	assert.True(t, isOwner("abcdefg", "/api/abcdefg/something", "/api/{uid}/something"))
	// Made to fail - should be case sensitive
	assert.False(t, isOwner("abcdefg", "/api/abcdefG/something", "/api/{uid}/something"))
	// Made to fail - should not turn true for non-uid
	assert.False(t, isOwner("abcdefg", "/api/abcdefg/something", "/api/{id}/something"))
	// When there is no placeholder the result should be false
	assert.False(t, isOwner("abcdefg", "/api/something/something", "/api/abcdefg/something"))
}
