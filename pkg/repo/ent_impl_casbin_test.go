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
		_, err = entImpl.AddPolicy(AuthzPolicy{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts/*",
			Action:  mergeAllowedHTTPMethods([]string{"*"}),
			Effect:  "allow",
		})
		assert.NoError(t, err)
		err = entImpl.enforcer.LoadPolicy()
		assert.NoError(t, err)

		shouldPass, err := entImpl.Enforce(AuthzRequest{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts/",
			Action:  "POST",
		})
		assert.True(t, shouldPass)
		assert.NoError(t, err)

		shouldPass, err = entImpl.Enforce(AuthzRequest{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts/somesubpath",
			Action:  "GET",
		})
		assert.True(t, shouldPass)
		assert.NoError(t, err)
	})

	t.Run("GET to /api/v1/facts/:id", func(t *testing.T) {
		hashedToken := uniuri.New()
		_, err = entImpl.AddPolicy(AuthzPolicy{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts/:id",
			Action:  mergeAllowedHTTPMethods([]string{"GET"}),
			Effect:  "allow",
		})
		assert.NoError(t, err)
		err = entImpl.enforcer.LoadPolicy()
		assert.NoError(t, err)

		shouldPass, err := entImpl.Enforce(AuthzRequest{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts/fact_somerandomid",
			Action:  "GET",
		})
		assert.True(t, shouldPass, "should pass because /api/v1/facts/fact_somerandomid matches the keyMatch2 function")
		assert.NoError(t, err)

		shouldPass, err = entImpl.Enforce(AuthzRequest{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts",
			Action:  "GET",
		})
		assert.False(t, shouldPass)
		assert.NoError(t, err, "should not pass because the GET path has less than what the policy defined")

		shouldPass, err = entImpl.Enforce(AuthzRequest{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts/fact_somerandomid/anything_else_in_the_path",
			Action:  "GET",
		})
		assert.False(t, shouldPass, "should not pass because the GET path has more than what the policy defined")
		assert.NoError(t, err)
	})

	t.Run("POST to /api/v1/facts", func(t *testing.T) {
		hashedToken := uniuri.New()
		_, err = entImpl.AddPolicy(AuthzPolicy{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts",
			Action:  mergeAllowedHTTPMethods([]string{"POST"}),
			Effect:  "allow",
		})
		assert.NoError(t, err)
		err = entImpl.enforcer.LoadPolicy()
		assert.NoError(t, err)

		shouldPass, err := entImpl.Enforce(AuthzRequest{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts",
			Action:  "POST",
		})
		assert.True(t, shouldPass)
		assert.NoError(t, err)

		shouldPass, err = entImpl.Enforce(AuthzRequest{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts",
			Action:  "GET",
		})
		assert.False(t, shouldPass)
		assert.NoError(t, err)
	})

	t.Run("POST to /api/v1/facts/noscope", func(t *testing.T) {
		hashedToken := uniuri.New()
		_, err = entImpl.AddPolicy(AuthzPolicy{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts/noscope",
			Action:  mergeAllowedHTTPMethods([]string{"POST"}),
			Effect:  "allow",
		})
		assert.NoError(t, err)
		err = entImpl.enforcer.LoadPolicy()
		assert.NoError(t, err)

		shouldPass, err := entImpl.Enforce(AuthzRequest{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts/noscope",
			Action:  "POST",
		})
		assert.True(t, shouldPass)
		assert.NoError(t, err)

		shouldPass, err = entImpl.Enforce(AuthzRequest{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts",
			Action:  "POST",
		})
		assert.False(t, shouldPass)
		assert.NoError(t, err)

		shouldPass, err = entImpl.Enforce(AuthzRequest{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts/",
			Action:  "POST",
		})
		assert.False(t, shouldPass)
		assert.NoError(t, err)

		shouldPass, err = entImpl.Enforce(AuthzRequest{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts/:id",
			Action:  "GET",
		})
		assert.False(t, shouldPass)
		assert.NoError(t, err)
	})

	t.Run("POST to /api/v1/facts/noscope with grouping policy - publishable_token", func(t *testing.T) {
		hashedToken := uniuri.New()
		group := "publishable_token"

		_, err = entImpl.AddPolicy(AuthzPolicy{
			Subject: group,
			Domain:  "example.com",
			Object:  "/api/v1/facts/noscope",
			Action:  mergeAllowedHTTPMethods([]string{"POST"}),
			Effect:  "allow",
		})
		assert.NoError(t, err)
		_, err = entImpl.AddGroupingPolicy(AuthzGroupingPolicy{
			Domain:  "example.com",
			Subject: hashedToken,
			Group:   group,
		})
		assert.NoError(t, err)
		err = entImpl.enforcer.LoadPolicy()
		assert.NoError(t, err)

		shouldPass, err := entImpl.Enforce(AuthzRequest{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts/noscope",
			Action:  "POST",
		})
		assert.True(t, shouldPass)
		assert.NoError(t, err)

		shouldPass, err = entImpl.Enforce(AuthzRequest{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts",
			Action:  "POST",
		})
		assert.False(t, shouldPass)
		assert.NoError(t, err)

		shouldPass, err = entImpl.Enforce(AuthzRequest{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts/",
			Action:  "POST",
		})
		assert.False(t, shouldPass)
		assert.NoError(t, err)

		shouldPass, err = entImpl.Enforce(AuthzRequest{
			Subject: hashedToken,
			Domain:  "example.com",
			Object:  "/api/v1/facts/:id",
			Action:  "GET",
		})
		assert.False(t, shouldPass)
		assert.NoError(t, err)
	})
}
