package cryptoutil

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

// Cipher represents a cipher text (encrypted message).
type Cipher []byte

// NewCipher creates a cipher from its hexadecimal string representation.
func NewCipher(hexString string) Cipher {
	c, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}

	return Cipher(c)
}

// String returns the hexadecimal representation of the cipher.
func (c Cipher) String() string {
	return hex.EncodeToString([]byte(c))
}

// Binary returns the string binary representation of the cipher's bytes.
func (c Cipher) Binary() string {
	var buffer bytes.Buffer
	for _, n := range c {
		buffer.WriteString(fmt.Sprintf("%08b", n))
	}

	return buffer.String()
}

// XOR does an exclusive or between two ciphers.
func (c Cipher) XOR(c2 Cipher) (Cipher, error) {
	b := make([]byte, max(len(c), len(c2)))
	for i := 0; i < len(b); i++ {
		if i < len(c) && i < len(c2) {
			b[i] = c[i] ^ c2[i]
			continue
		}

		if i < len(c) {
			b[i] = c[i]
		} else {
			b[i] = c2[i]
		}
	}

	return Cipher(b), nil
}

func max(i, j int) int {
	if i < j {
		return j
	}

	return i
}
