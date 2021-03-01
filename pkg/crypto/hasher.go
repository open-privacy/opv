package crypto

import (
	"encoding/base64"
	"fmt"

	"github.com/open-privacy/opv/pkg/config"
	"golang.org/x/crypto/scrypt"
	"golang.org/x/crypto/sha3"
)

const (
	hasherScrypt    = "scrypt"
	hasherKeccak256 = "keccak256"
)

// Hasher is an interface
type Hasher interface {
	Hash(string) string
	HashFaster(string) string
}

// MustNewHasher creates a new Hasher
func MustNewHasher() Hasher {
	var hasher Hasher
	switch config.ENV.HasherName {
	case hasherScrypt:
		hasher = &ScryptHasher{
			salt:   []byte(config.ENV.HasherScryptSalt),
			n:      config.ENV.HasherScryptN,
			r:      8,
			p:      1,
			keyLen: 32,
		}
	case hasherKeccak256:
		hasher = &Keccak256Hasher{}
	default:
		panic(fmt.Sprintf("unknown hasher name: %s", config.ENV.HasherName))
	}
	return hasher
}

// Keccak256Hasher is a Hasher, it was proven to be secure and used in Ethereum
type Keccak256Hasher struct{}

// Hash ...
func (k *Keccak256Hasher) Hash(s string) string {
	h := make([]byte, 64)
	sha3.ShakeSum256(h, []byte(s))
	return string(h)
}

// HashFaster ...
func (k *Keccak256Hasher) HashFaster(s string) string {
	return k.Hash(s)
}

// ScryptHasher is a Hasher that implements the Scrypt hashing algorithm
type ScryptHasher struct {
	salt   []byte
	n      int
	r      int
	p      int
	keyLen int
}

// Hash ...
func (sh *ScryptHasher) Hash(s string) string {
	dk, _ := scrypt.Key(
		[]byte(s),
		sh.salt,
		sh.n,
		sh.r,
		sh.p,
		sh.keyLen,
	)
	return base64.StdEncoding.EncodeToString(dk)
}

// HashFaster ...
func (sh *ScryptHasher) HashFaster(s string) string {
	dk, _ := scrypt.Key(
		[]byte(s),
		sh.salt,
		sh.n>>4,
		sh.r,
		sh.p,
		sh.keyLen,
	)
	return base64.StdEncoding.EncodeToString(dk)
}
