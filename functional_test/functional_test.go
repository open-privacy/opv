package functional_test

import (
	"net/http"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
	"github.com/avast/retry-go"
	"github.com/caarlos0/env/v6"
	"github.com/dchest/uniuri"
	"github.com/open-privacy/opv/pkg/config"
	"github.com/stretchr/testify/assert"
)

func init() {
	if err := env.Parse(&TESTENV); err != nil {
		panic(err)
	}
}

// TESTENV is the env configuration for functional testing
var TESTENV = struct {
	ControlplaneHostport string `env:"TESTENV_CONTROLPLANE_HOSTPORT" envDefault:"http://127.0.0.1:27999"`
	DataplaneHostport    string `env:"TESTENV_DATAPLANE_HOSTPORT" envDefault:"http://127.0.0.1:28000"`
	DefaultDomain        string `env:"TESTENV_DEFAULT_DOMAIN" envDefault:"example.com"`
}{}

var assertGetToken = func(t *testing.T, allowedHttpMethods []string) string {
	var token string
	Test(
		t,
		Description("Post to controlplane to create a grant"),
		Post(TESTENV.ControlplaneHostport+"/api/v1/grants"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().JSON(map[string]interface{}{
			"domain":               TESTENV.DefaultDomain,
			"allowed_http_methods": allowedHttpMethods,
		}),

		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".token").Len().GreaterThan(0),
		Store().Response().Body().JSON().JQ(".token").In(&token),
	)
	time.Sleep(config.ENV.AuthzCasbinAutoloadInterval + time.Second)
	return token
}

func TestHealthz(t *testing.T) {
	var err error

	err = retry.Do(func() error {
		return Do(
			Description("Wait for healthz check to be passed"),
			Get(TESTENV.ControlplaneHostport+"/api/v1/healthz"),
			Send().Headers("Content-Type").Add("application/json"),
			Expect().Status().Equal(http.StatusOK),
		)
	})
	assert.NoError(t, err)

	err = retry.Do(func() error {
		return Do(
			Description("Wait for healthz check to be passed"),
			Get(TESTENV.DataplaneHostport+"/api/v1/healthz"),
			Send().Headers("Content-Type").Add("application/json"),
			Expect().Status().Equal(http.StatusOK),
		)
	})
	assert.NoError(t, err)
}

func TestCreateGrant(t *testing.T) {
	Test(
		t,
		Description("Happy code path: Post to controlplane to create a grant"),
		Post(TESTENV.ControlplaneHostport+"/api/v1/grants"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().JSON(map[string]interface{}{
			"domain":               TESTENV.DefaultDomain,
			"allowed_http_methods": []string{"*"},
		}),

		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".token").Len().GreaterThan(0),
	)

	Test(
		t,
		Description("Error due to unexpected HTTP method"),
		Post(TESTENV.ControlplaneHostport+"/api/v1/grants"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().JSON(map[string]interface{}{
			"domain":               TESTENV.DefaultDomain,
			"allowed_http_methods": []string{"INVALID"},
		}),

		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".code").NotEqual(0),
	)

	Test(
		t,
		Description("Error due to unexpected domain"),
		Post(TESTENV.ControlplaneHostport+"/api/v1/grants"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().JSON(map[string]interface{}{
			"domain":               "*.invalid.domain",
			"allowed_http_methods": []string{"*"},
		}),

		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".code").NotEqual(0),
	)
}

func TestCreateFact(t *testing.T) {
	token := assertGetToken(t, []string{"POST"})

	t.Run("happy code path", func(t *testing.T) {
		scopeID := uniuri.NewLen(uniuri.UUIDLen)
		factTypeSlug := "ssn"

		Test(
			t,
			Description("Post to dataplane to create a fact"),
			Post(TESTENV.DataplaneHostport+"/api/v1/facts"),
			Send().Headers("Content-Type").Add("application/json"),
			Send().Headers("X-OPV-GRANT-TOKEN").Add(token),
			Send().Body().JSON(map[string]interface{}{
				"scope_custom_id": scopeID,
				"fact_type_slug":  factTypeSlug,
				"value":           "123-45-6789",
			}),

			Expect().Status().Equal(http.StatusOK),
			Expect().Body().JSON().JQ(".id").NotEqual(""),
			Expect().Body().JSON().JQ(".scope_custom_id").Equal(scopeID),
			Expect().Body().JSON().JQ(".fact_type_slug").Equal(factTypeSlug),
		)
	})

	t.Run("ssn fact type slug", func(t *testing.T) {
		t.Run("valid ssns", func(t *testing.T) {
			scopeID := uniuri.NewLen(uniuri.UUIDLen)
			factTypeSlug := "ssn"
			validSSNs := []string{
				"123-45-6789",
				"123456789",
				"123 45 6789",
			}

			for _, ssn := range validSSNs {
				Test(
					t,
					Description("Post to dataplane to create a fact"),
					Post(TESTENV.DataplaneHostport+"/api/v1/facts"),
					Send().Headers("Content-Type").Add("application/json"),
					Send().Headers("X-OPV-GRANT-TOKEN").Add(token),
					Send().Body().JSON(map[string]interface{}{
						"scope_custom_id": scopeID,
						"fact_type_slug":  factTypeSlug,
						"value":           ssn,
					}),

					Expect().Status().Equal(http.StatusOK),
					Expect().Body().JSON().JQ(".id").NotEqual(""),
				)
			}
		})

		t.Run("error with invalid ssn", func(t *testing.T) {
			scopeID := uniuri.NewLen(uniuri.UUIDLen)
			factTypeSlug := "ssn"
			invalidSSNs := []string{
				"invalid",
				"1234",
			}

			for _, ssn := range invalidSSNs {
				Test(
					t,
					Description("Post to dataplane to create a fact"),
					Post(TESTENV.DataplaneHostport+"/api/v1/facts"),
					Send().Headers("Content-Type").Add("application/json"),
					Send().Headers("X-OPV-GRANT-TOKEN").Add(token),
					Send().Body().JSON(map[string]interface{}{
						"scope_custom_id": scopeID,
						"fact_type_slug":  factTypeSlug,
						"value":           ssn,
					}),

					Expect().Status().Equal(http.StatusBadRequest),
				)
			}
		})
	})
}

func TestMalformattedJSON(t *testing.T) {
	token := assertGetToken(t, []string{"POST"})

	Test(
		t,
		Description("Post to dataplane with malformatted JSON"),
		Post(TESTENV.DataplaneHostport+"/api/v1/facts"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Headers("X-OPV-GRANT-TOKEN").Add(token),
		Send().Body().JSON("{"),

		Expect().Status().Equal(http.StatusBadRequest),
	)
}
