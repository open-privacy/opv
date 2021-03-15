package repo

import (
	"testing"

	"github.com/dchest/uniuri"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestEntImplCasbinRule(t *testing.T) {
	entImpl, err := newEntImpl(log.New("test"))
	assert.NoError(t, err)

	t.Run("* actions", func(t *testing.T) {
		hashedToken := uniuri.New()
		_, err = entImpl.enforcer.AddPolicy(
			hashedToken,
			"example.com",
			"/api/v1/facts",
			mergeAllowedHTTPMethods([]string{"*"}),
			"allow",
		)
		assert.NoError(t, err)
		err = entImpl.enforcer.LoadPolicy()
		assert.NoError(t, err)

		shouldPass, err := entImpl.enforcer.Enforce(hashedToken, "example.com", "/api/v1/facts", "POST")
		assert.True(t, shouldPass)
		assert.NoError(t, err)
	})

	t.Run("POST to /api/v1/facts", func(t *testing.T) {
		hashedToken := uniuri.New()
		_, err = entImpl.enforcer.AddPolicy(
			hashedToken,
			"example.com",
			"/api/v1/facts",
			mergeAllowedHTTPMethods([]string{"POST"}),
			"allow",
		)
		assert.NoError(t, err)
		err = entImpl.enforcer.LoadPolicy()
		assert.NoError(t, err)

		shouldPass, err := entImpl.enforcer.Enforce(hashedToken, "example.com", "/api/v1/facts", "POST")
		assert.True(t, shouldPass)
		assert.NoError(t, err)

		shouldPass, err = entImpl.enforcer.Enforce(hashedToken, "example.com", "/api/v1/facts/:id", "GET")
		assert.False(t, shouldPass)
		assert.NoError(t, err)
	})

	t.Run("POST to /api/v1/facts/noscope", func(t *testing.T) {
		hashedToken := uniuri.New()
		_, err = entImpl.enforcer.AddPolicy(
			hashedToken,
			"example.com",
			"/api/v1/facts/noscope",
			mergeAllowedHTTPMethods([]string{"POST"}),
			"allow",
		)
		assert.NoError(t, err)
		err = entImpl.enforcer.LoadPolicy()
		assert.NoError(t, err)

		shouldPass, err := entImpl.enforcer.Enforce(hashedToken, "example.com", "/api/v1/facts/noscope", "POST")
		assert.True(t, shouldPass)
		assert.NoError(t, err)

		shouldPass, err = entImpl.enforcer.Enforce(hashedToken, "example.com", "/api/v1/facts", "POST")
		assert.False(t, shouldPass)
		assert.NoError(t, err)

		shouldPass, err = entImpl.enforcer.Enforce(hashedToken, "example.com", "/api/v1/facts/", "POST")
		assert.False(t, shouldPass)
		assert.NoError(t, err)

		shouldPass, err = entImpl.enforcer.Enforce(hashedToken, "example.com", "/api/v1/facts/:id", "GET")
		assert.False(t, shouldPass)
		assert.NoError(t, err)
	})
}
