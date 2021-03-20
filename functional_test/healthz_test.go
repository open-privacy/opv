package functional_test

import (
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
	"github.com/avast/retry-go"
	"github.com/stretchr/testify/assert"
)

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
