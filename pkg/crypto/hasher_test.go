package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeccak256Hasher(t *testing.T) {
	h := &Keccak256Hasher{}
	data, salt := "data", "salt"
	hashed := h.Hash(data, salt)
	assert.Equal(t, "YJN04bmz54UvvaISrD/SXTkicmpxv1NRyxYbXdU5hZaCde6EpOA+bOFPgidvlD3MJrLozC2aUpqkFiO7LUWN2Q==", hashed)
}
