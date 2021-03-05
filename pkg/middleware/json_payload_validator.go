package middleware

import(
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/apimodel"
)

func ValidateJSONPayload(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var bodyBytes []byte
		if c.Request().Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
			// Restore the io.ReadCloser to its original state
			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			// Use the content
			if len(bodyBytes) > 0 {
				jsonBody := make(map[string]interface{})
				err := json.Unmarshal(bodyBytes, &jsonBody)

				if err != nil && err != io.EOF {
					return apimodel.FormatHTTPError(c, apimodel.ErrJSONMalformatted)
				}
			}
		}

		return next(c)
	}
}
