package crypto

import (
	"encoding/base64"
	"fmt"

	"github.com/open-privacy/opv/pkg/config"
	"golang.org/x/crypto/scrypt"
)

const (
	hasherScrypt = "scrypt"
)

type Hasher interface {
	Hash(string) string
	HashFaster(string) string
}

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
	default:
		panic(fmt.Sprintf("unknown hasher name: %s", config.ENV.HasherName))
	}
	return hasher
}

type ScryptHasher struct {
	salt   []byte
	n      int
	r      int
	p      int
	keyLen int
}

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
