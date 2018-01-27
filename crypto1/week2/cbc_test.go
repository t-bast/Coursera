package cbc

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPadding(t *testing.T) {
	t.Run("Pad incomplete block", func(t *testing.T) {
		message := []byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
		assert.Len(t, message, BlockSize-3)

		padded := pad(message)
		assert.Len(t, padded, BlockSize)
		assert.EqualValues(t,
			[]byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3},
			padded)
	})

	t.Run("Add new block if necessary", func(t *testing.T) {
		message := []byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
		assert.Len(t, message, BlockSize)
		padded := pad(message)
		assert.Len(t, padded, 2*BlockSize)

		for i := 0; i < BlockSize; i++ {
			assert.Equal(t, byte(2), padded[i])
			assert.Equal(t, byte(BlockSize), padded[BlockSize+i])
		}
	})
}

func TestEncryptLength(t *testing.T) {
	iv, err := generateRandomBytes(BlockSize)
	require.NoError(t, err)
	key, err := generateRandomBytes(BlockSize)
	require.NoError(t, err)

	t.Run("Pad last block", func(t *testing.T) {
		message := []byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
		assert.True(t, len(message) < BlockSize)

		c, err := Encrypt(iv, key, message)
		assert.NoError(t, err)
		assert.Len(t, c, 2*BlockSize)
	})

	t.Run("Add last padding block if needed", func(t *testing.T) {
		message := []byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
		assert.Len(t, message, BlockSize)

		c, err := Encrypt(iv, key, message)
		assert.NoError(t, err)
		assert.Len(t, c, 3*BlockSize)
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
