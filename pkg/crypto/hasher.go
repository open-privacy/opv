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
	Hash(data, salt string) string
	HashFaster(data, salt string) string
}

// MustNewHasher creates a new Hasher
func MustNewHasher() Hasher {
	var hasher Hasher
	switch config.ENV.HasherName {
	case hasherScrypt:
		hasher = &ScryptHasher{
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
func (k *Keccak256Hasher) Hash(data, salt string) string {
	h := make([]byte, 64)
	sha3.ShakeSum256(h, []byte(salt+data))
	return string(h)
}

// HashFaster ...
func (k *Keccak256Hasher) HashFaster(data, salt string) string {
	return k.Hash(data, salt)
}

// ScryptHasher is a Hasher that implements the Scrypt hashing algorithm
type ScryptHasher struct {
	n      int
	r      int
	p      int
	keyLen int
}

// Hash ...
func (sh *ScryptHasher) Hash(data, salt string) string {
	dk, _ := scrypt.Key(
		[]byte(data),
		[]byte(salt),
		sh.n,
		sh.r,
		sh.p,
		sh.keyLen,
	)
	return base64.StdEncoding.EncodeToString(dk)
}

// HashFaster ...
func (sh *ScryptHasher) HashFaster(data, salt string) string {
	dk, _ := scrypt.Key(
		[]byte(data),
		[]byte(salt),
		sh.n>>4,
		sh.r,
		sh.p,
		sh.keyLen,
	)
	return base64.StdEncoding.EncodeToString(dk)
}
