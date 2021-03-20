package functional_test

import (
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
)

func TestProxyPlane(t *testing.T) {
	t.Run("tokenize and then detokenize", func(t *testing.T) {
		ssn := "123-45-6789"
		ssnFactID := ""

		Test(t,
			Description("Post to proxy to httpbin.org - 200 for tokenize"),
			Post(TESTENV.ProxyplaneHostport+"/tokenize"),
			Send().Headers("Content-Type").Add("application/json"),
			Send().Body().JSON(map[string]interface{}{
				"user": map[string]interface{}{
					"ssn": ssn,
				},
			}),

			Expect().Status().Equal(http.StatusOK),
			Expect().Body().JSON().JQ(".json.user.ssn").Contains("fact_"),
			Store().Response().Body().JSON().JQ(".json.user.ssn").In(&ssnFactID),
		)

		Test(t,
			Description("Post to proxy to httpbin.org - 200 for detokenize"),
			Post(TESTENV.ProxyplaneHostport+"/detokenize"),
			Send().Headers("Content-Type").Add("application/json"),
			Send().Body().JSON(map[string]interface{}{
				"user": map[string]interface{}{
					"ssn": ssnFactID,
				},
			}),

			Expect().Status().Equal(http.StatusOK),
			Expect().Body().JSON().JQ(".json.user.ssn").Equal(ssn),
		)
	})
}
