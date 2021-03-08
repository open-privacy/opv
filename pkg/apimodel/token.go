package apimodel

import (
	"fmt"
	"strings"

	"github.com/dchest/uniuri"
	"github.com/open-privacy/opv/pkg/crypto"
)

// GrantToken represents a secret token
type GrantToken struct {
	Version string
	Domain  string

	secret string
}

// String returns the plaintext encoding of the Token
func (t *GrantToken) String() string {
	return fmt.Sprintf("%s:%s:%s", t.Version, t.Domain, t.secret)
}

// Hash uses HashFaster to hash the grant
func (t *GrantToken) Hash(h crypto.Hasher) string {
	return h.HashFaster(t.String(), t.Domain)
}

// NewToken creates a new Token
func NewToken(version string, domain string) (*GrantToken, error) {
	var t *GrantToken

	switch version {
	case "v1":
		secret := uniuri.NewLen(uniuri.UUIDLen)
		t = &GrantToken{
			Version: version,
			Domain:  domain,
			secret:  secret,
		}
	default:
		return nil, fmt.Errorf("failed to run NewToken, invalid version %s", version)
	}

	return t, nil
}

// ParseFromString creates a new Token from its plaintext string
func (t *GrantToken) ParseFromString(s string) error {
	parts := strings.Split(s, ":")

	if len(parts) < 1 {
		return fmt.Errorf("invalid token")
	}

	version := parts[0]
	domain := ""
	secret := ""

	switch version {
	case "v1":
		if len(parts) != 3 {
			return fmt.Errorf("failed to run ParseDomain, invalid token")
		}
		domain = parts[1]
		secret = parts[2]
	default:
		return fmt.Errorf("failed to run ParseDomain, invalid version %s", version)
	}

	t.Version = version
	t.Domain = domain
	t.secret = secret
	return nil
}
