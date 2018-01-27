package ctr

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	res0 := add(iv, 0)

	assert.EqualValues(t, iv, res0)

	res1 := add(iv, 1)
	assert.EqualValues(t,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01},
		res1)

	res := add(iv, 1<<20)
	assert.EqualValues(t,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x10, 0, 0},
		res)
}

func TestXor(t *testing.T) {
	b1 := []byte{0x32, 0xa7, 0x42}
	b2 := []byte{0x20, 0x12}
	res1 := xor(b1, b2)
	res2 := xor(b2, b1)

	assert.Len(t, res1, 2)
	assert.EqualValues(t, res1, res2)
	assert.EqualValues(t, []byte{0x12, 0xb5}, res2)
}

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
		message := []byte("spongebob")
		c, err := Encrypt(iv, key, message)
		assert.NoError(t, err)

		d, err := Decrypt(key, c)
		assert.NoError(t, err)

		assert.Equal(t, "spongebob", string(d))
	})

	t.Run("Encrypt and decrypt long message", func(t *testing.T) {
		message := []byte("hello mr spongebob")
		c, err := Encrypt(iv, key, message)
		assert.NoError(t, err)

		d, err := Decrypt(key, c)
		assert.NoError(t, err)

		assert.Equal(t, "hello mr spongebob", string(d))
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
