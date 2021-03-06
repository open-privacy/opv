package modifier

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/hashicorp/go-retryablehttp"
)

const (
	headerOPVGrantToken  = "x-opv-grant-token"
	headerOPVUserAgent   = "OPV-ProxyPlane"
	headerOPVContentType = "application/json"
)

var (
	client     *retryablehttp.Client
	clientOnce sync.Once
)

type conn struct {
	dpGrantToken string
	dpURL        string
	client       *retryablehttp.Client
}

func newConn(dpGrantToken string, dpURL string) *conn {
	clientOnce.Do(func() {
		client = retryablehttp.NewClient()
	})

	return &conn{
		dpGrantToken: dpGrantToken,
		dpURL:        dpURL,
		client:       client,
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
	req.Header.Add("Content-Type", headerOPVContentType)
	req.Header.Add("User-Agent", headerOPVUserAgent)

	resp, err := c.client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	}
	if resp.StatusCode >= 400 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("createFact api failed %s", string(respBody))
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
	req.Header.Add("User-Agent", headerOPVUserAgent)

	resp, err := c.client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	}
	if resp.StatusCode >= 400 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("getFact api failed %s", string(respBody))
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
