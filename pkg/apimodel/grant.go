package apimodel

import (
	"fmt"
	"strings"

	"github.com/dchest/uniuri"
	"github.com/open-privacy/opv/pkg/crypto"
)

// CreateGrant represents the request object of creating a grant
type CreateGrant struct {
	Domain         string   `json:"domain"`
	AllowedActions []string `json:"allowed_actions"`

	// TODO we can add more predicates here to limit how we want to expose
	// objects that the grant can operate on
}

// Grant represents the grant object
type Grant struct {
	Token          string   `json:"token"`
	AllowedActions []string `json:"allowed_actions"`
}

// GenToken sets a token for the grant
func (g *Grant) GenToken(version string, domain string) {
	switch version {
	case "v1":
		g.Token = fmt.Sprintf("v1:%s:%s",
			domain,
			uniuri.NewLen(uniuri.UUIDLen),
		)
	}
	return
}

// Hash uses HashFaster to hash the grant
func (g *Grant) Hash(h crypto.Hasher) string {
	return h.HashFaster(g.Token)
}

// Domain gets the domain from the token of the grant
func (g *Grant) Domain() (string, error) {
	parts := strings.Split(g.Token, ":")

	if len(parts) < 1 {
		return "", fmt.Errorf("invalid token")
	}

	version := parts[0]

	switch version {
	case "v1":
		if len(parts) != 3 {
			return "", fmt.Errorf("invalid token")
		}
		return parts[1], nil
	}

	return "", fmt.Errorf("invalid version: %v", version)
}
