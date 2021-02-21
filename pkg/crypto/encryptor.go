package crypto

import (
	"fmt"

	"golang.org/x/crypto/nacl/secretbox"

	"github.com/open-privacy/opv/pkg/config"
)

const (
	encryptorSecretbox      = "secretbox"
	encryptorHashicorpVault = "hashicorp_vault"
)

type Encryptor interface {
	Encrypt(nonce, plaintext string) (ciphertext string, err error)
	Decrypt(nonce, ciphertext string) (plaintext string, err error)
}

type SecretboxEncryptor struct {
	keys [][32]byte
}

func (se *SecretboxEncryptor) Encrypt(nonce, plaintext string) (ciphertext string, err error) {
	nonceBytes := [24]byte{}
	copy(nonceBytes[:], []byte(nonce))

	var out []byte
	out = secretbox.Seal(out[:0], []byte(plaintext), &nonceBytes, &se.keys[0])
	return string(out), nil
}

func (se *SecretboxEncryptor) Decrypt(nonce, ciphertext string) (plaintext string, err error) {
	nonceBytes := [24]byte{}
	copy(nonceBytes[:], []byte(nonce))

	for i := range se.keys {
		var opened []byte
		opened, ok := secretbox.Open(opened[:0], []byte(ciphertext), &nonceBytes, &se.keys[i])
		if ok {
			return string(opened), nil
		}
	}

	return "", fmt.Errorf("ciphertext could not be decrypted")
}

func MustNewEncryptor() Encryptor {
	var encryptor Encryptor
	switch config.ENV.EncryptorName {
	case encryptorSecretbox:
		keys := [][32]byte{}
		for _, k := range config.ENV.EncryptorSecretboxKeys {
			var b [32]byte
			copy(b[:], []byte(k))
			keys = append(keys, b)
		}
		encryptor = &SecretboxEncryptor{keys: keys}
	default:
		panic(fmt.Sprintf("unknown encryptor name: %s", config.ENV.EncryptorName))
	}
	return encryptor
}
