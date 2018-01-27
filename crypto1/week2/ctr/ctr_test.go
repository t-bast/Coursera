package ctr

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	iv, err := generateRandomBytes(BlockSize)
	require.NoError(t, err)
	key, err := generateRandomBytes(BlockSize)
	require.NoError(t, err)

	t.Run("First encrypted block is IV", func(t *testing.T) {
		message := []byte("hello")
		c, err := Encrypt(iv, key, message)
		assert.NoError(t, err)

		for i := 0; i < BlockSize; i++ {
			assert.Equal(t, iv[i], c[i])
		}
	})

	t.Run("No padding is necessary", func(t *testing.T) {
		message := []byte("hello")
		c, err := Encrypt(iv, key, message)
		assert.NoError(t, err)
		assert.Len(t, c, BlockSize+5)
	})

	t.Run("Encrypt and decrypt message", func(t *testing.T) {
		assert.Fail(t, "Not implemented")
	})
}

// generateRandomBytes generate a random byte array
// with a given number of bytes.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
