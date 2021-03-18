package modifier

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/open-privacy/opv/pkg/config"
)

const headerOPVGrantToken = "x-opv-grant-token"

type conn struct {
	dpGrantToken string
	dpURL        string
	client       *retryablehttp.Client
}

func mustNewConn() *conn {
	if config.ENV.ProxyPlaneDPGrantToken == "" {
		panic("OPV_PROXY_PLANE_DP_GRANT_TOKEN should not be empty")
	}

	return &conn{
		dpGrantToken: config.ENV.ProxyPlaneDPGrantToken,
		dpURL:        config.ENV.ProxyPlaneDPURL,
		client:       retryablehttp.NewClient(),
	}
}

func (c *conn) createFact(factTypeSlug string, value string) (factID string, err error) {
	req, err := retryablehttp.NewRequest(
		"POST",
		c.dpURL+"/api/v1/facts",
		strings.NewReader(
			fmt.Sprintf(`{"fact_type_slug":"%s", "value": %s}`, factTypeSlug, value),
		),
	)
	if err != nil {
		return "", err
	}
	req.Header.Add(headerOPVGrantToken, c.dpGrantToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	}

	f := struct {
		ID string `json:"id"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&f)
	if err != nil {
		return "", err
	}

	return f.ID, nil
}

func (c *conn) getFact(factID string) (value string, err error) {
	req, err := retryablehttp.NewRequest(
		"GET",
		c.dpURL+"/api/v1/facts/"+factID,
		strings.NewReader(""),
	)
	if err != nil {
		return "", err
	}
	req.Header.Add(headerOPVGrantToken, c.dpGrantToken)

	resp, err := c.client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	}

	f := struct {
		Value string `json:"value"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&f)
	if err != nil {
		return "", err
	}

	return f.Value, nil
}
