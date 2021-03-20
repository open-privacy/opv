package functional_test

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
	"github.com/dchest/uniuri"
	"github.com/open-privacy/opv/pkg/config"
)

func generateScopeID() string {
	return uniuri.NewLen(uniuri.UUIDLen)
}

var getValidTokenMemo = make(map[string]string)
var getValidToken = func(t *testing.T, allowedHttpMethods []string, paths []string) string {
	cacheKey := strings.Join(
		append(allowedHttpMethods, paths...),
		"",
	)
	token, ok := getValidTokenMemo[cacheKey]
	if ok {
		return token
	}

	Test(
		t,
		Description("Post to controlplane to create a grant"),
		Post(TESTENV.ControlplaneHostport+"/api/v1/grants"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().JSON(map[string]interface{}{
			"domain":               TESTENV.DefaultDomain,
			"allowed_http_methods": allowedHttpMethods,
			"paths":                paths,
		}),

		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".token").Len().GreaterThan(0),
		Store().Response().Body().JSON().JQ(".token").In(&token),
	)
	time.Sleep(config.ENV.AuthzCasbinAutoloadInterval + time.Second)

	getValidTokenMemo[cacheKey] = token
	return token
}

var assertCreateFact = func(t *testing.T, token, scopeID string, factValue string) string {
	var factID string

	Test(
		t,
		Description("Post to dataplane to create a fact"),
		Post(TESTENV.DataplaneHostport+"/api/v1/facts"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Headers("X-OPV-GRANT-TOKEN").Add(token),
		Send().Body().JSON(map[string]interface{}{
			"scope_custom_id": scopeID,
			"value":           factValue,
			"fact_type_slug":  "ascii",
		}),

		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".id").Len().GreaterThan(0),
		Store().Response().Body().JSON().JQ(".id").In(&factID),
	)
	return factID
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
	token := getValidToken(t, []string{"POST"}, nil)

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

	t.Run("test not supported fact type slug", func(t *testing.T) {
		scopeID := uniuri.NewLen(uniuri.UUIDLen)
		factTypeSlug := "invalid_slug"

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

			Expect().Status().Equal(http.StatusBadRequest),
		)
	})
}

func TestCreateFactUniqueScopeConstraint(t *testing.T) {
	token := getValidToken(t, []string{"POST"}, nil)

	t.Run("happy code path", func(t *testing.T) {
		scopeID := generateScopeID()
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

	t.Run("happy code path: create facts with empty scope with the same value multiple times", func(t *testing.T) {
		n := 2

		for i := 0; i < n; i++ {
			Test(
				t,
				Description("Post to dataplane to create a fact"),
				Post(TESTENV.DataplaneHostport+"/api/v1/facts"),
				Send().Headers("Content-Type").Add("application/json"),
				Send().Headers("X-OPV-GRANT-TOKEN").Add(token),
				Send().Body().JSON(map[string]interface{}{
					// we don't pass scope_custom_id here, so these are associated with the same "empty" scope
					"fact_type_slug": "ssn",
					"value":          "123-45-6789",
				}),

				Expect().Status().Equal(http.StatusOK),
				Expect().Body().JSON().JQ(".id").NotEqual(""),
				Expect().Body().JSON().JQ(".scope_custom_id").Equal(nil),
				Expect().Body().JSON().JQ(".fact_type_slug").Equal("ssn"),
			)
		}
	})

	t.Run("error: create facts with non-empty scope with the same value multiple times", func(t *testing.T) {
		scopeCustomID := uniuri.NewLen(uniuri.UUIDLen)

		// first time should work
		Test(
			t,
			Description("Post to dataplane to create a fact"),
			Post(TESTENV.DataplaneHostport+"/api/v1/facts"),
			Send().Headers("Content-Type").Add("application/json"),
			Send().Headers("X-OPV-GRANT-TOKEN").Add(token),
			Send().Body().JSON(map[string]interface{}{
				"scope_custom_id": scopeCustomID,
				"fact_type_slug":  "ssn",
				"value":           "123-45-6789",
			}),

			Expect().Status().Equal(http.StatusOK),
		)

		// second time should fail
		Test(
			t,
			Description("Post to dataplane to create a fact"),
			Post(TESTENV.DataplaneHostport+"/api/v1/facts"),
			Send().Headers("Content-Type").Add("application/json"),
			Send().Headers("X-OPV-GRANT-TOKEN").Add(token),
			Send().Body().JSON(map[string]interface{}{
				"scope_custom_id": scopeCustomID,
				"fact_type_slug":  "ssn",
				"value":           "123-45-6789",
			}),

			Expect().Status().NotEqual(http.StatusOK),
		)
	})
}

func TestCreateFactWithSlugValidation(t *testing.T) {
	token := getValidToken(t, []string{"POST"}, nil)

	t.Run("ssn fact type slug", func(t *testing.T) {
		t.Run("valid ssns", func(t *testing.T) {
			scopeID := generateScopeID()
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
			scopeID := generateScopeID()
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

func TestCreateFactFromJSV1(t *testing.T) {
	token := getValidToken(t, []string{"POST"}, []string{"/js/v1/facts"})

	t.Run("happy code path", func(t *testing.T) {
		scopeID := uniuri.NewLen(uniuri.UUIDLen)
		factTypeSlug := "ssn"

		Test(
			t,
			Description("OK to Post to dataplane to create a fact as a js publishable token"),
			Post(TESTENV.DataplaneHostport+"/js/v1/facts"),
			Send().Headers("Content-Type").Add("application/json"),
			Send().Headers("X-OPV-GRANT-TOKEN").Add(token),
			Send().Body().JSON(map[string]interface{}{
				"scope_custom_id": scopeID,
				"fact_type_slug":  factTypeSlug,
				"value":           "123-45-6789",
			}),

			Expect().Status().Equal(http.StatusOK),
			Expect().Body().JSON().JQ(".id").NotEqual(""),
			Expect().Body().JSON().JQ(".fact_type_slug").Equal(factTypeSlug),

			// scope custom id will be nil from POST /js/v1/facts
			Expect().Body().JSON().JQ(".scope_custom_id").Equal(nil),
		)

		Test(
			t,
			Description("Not OK to Post to apiv1"),
			Post(TESTENV.DataplaneHostport+"/api/v1/facts"),
			Send().Headers("Content-Type").Add("application/json"),
			Send().Headers("X-OPV-GRANT-TOKEN").Add(token),
			Send().Body().JSON(map[string]interface{}{
				"scope_custom_id": scopeID,
				"fact_type_slug":  factTypeSlug,
				"value":           "123-45-6789",
			}),

			Expect().Status().Equal(http.StatusUnauthorized),
		)
	})
}

func TestGetFact(t *testing.T) {
	token := getValidToken(t, []string{"POST", "GET"}, nil)
	factValue := fmt.Sprintf("%d%s", time.Now().UnixNano(), "_secret")
	scopeID := generateScopeID()
	factID := assertCreateFact(t, token, scopeID, factValue)

	t.Run("happy code path", func(t *testing.T) {
		Test(
			t,
			Description("GET to dataplane to retrieve a fact"),
			Get(TESTENV.DataplaneHostport+"/api/v1/facts/"+factID),
			Send().Headers("Content-Type").Add("application/json"),
			Send().Headers("X-OPV-GRANT-TOKEN").Add(token),

			Expect().Status().Equal(http.StatusOK),
			Expect().Body().JSON().JQ(".id").Equal(factID),
		)
	})
}

func TestMalformattedJSON(t *testing.T) {
	token := getValidToken(t, []string{"POST"}, nil)

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

func TestAPIAuditLogs(t *testing.T) {
	t.Run("it should return the correct API audit logs", func(t *testing.T) {
		Test(
			t,
			Description("send a health check request"),
			Get(TESTENV.DataplaneHostport+"/api/v1/healthz"),
			Send().Headers("Content-Type").Add("application/json"),
			Expect().Status().Equal(http.StatusOK),
		)

		Test(
			t,
			Description("should return the correct audit logs"),
			Get(
				fmt.Sprintf(
					"%s/api/v1/api_audits?domain=%s&path=%s",
					TESTENV.ControlplaneHostport,
					TESTENV.DefaultDomain,
					url.PathEscape("/api/v1/healthz"),
				),
			),
			Send().Headers("Content-Type").Add("application/json"),
			Expect().Status().Equal(http.StatusOK),
			Expect().Body().JSON().Len().GreaterThan(0),
		)

		Test(
			t,
			Description("should return the correct audit logs with limit and offset"),
			Get(
				fmt.Sprintf(
					"%s/api/v1/api_audits?domain=%s&path=%s&limit=2&offset=1&order_by=created_at&order_desc=true",
					TESTENV.ControlplaneHostport,
					TESTENV.DefaultDomain,
					url.PathEscape("/api/v1/healthz"),
				),
			),
			Send().Headers("Content-Type").Add("application/json"),
			Expect().Status().Equal(http.StatusOK),
			Expect().Body().JSON().Len().Equal(2),
		)
	})
}
