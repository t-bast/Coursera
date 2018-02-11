package week4

import (
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/t-bast/coursera/crypto1/cryptoutil"
)

// IV extracts the IV from the encrypted cipher text.
func IV() []byte {
	c := cryptoutil.NewCipher(EncryptedCipherText)
	return c[:BlockSize]
}

// BlockCount returns the number of blocks in the cipher text.
// Since the cipher text is valid, the number of bytes is exactly
// a multiple of the BlockSize.
func BlockCount() int {
	c := cryptoutil.NewCipher(EncryptedCipherText)
	return len(c) / BlockSize
}

// Decrypt uses the random oracle attack to decrypt the given
// encrypted cipher text.
func Decrypt() []byte {
	var decrypted []byte
	for blockNumber := 0; blockNumber < BlockCount()-1; blockNumber++ {
		fmt.Printf("Decrypting block %d...\n", blockNumber)
		decryptedBlock := DecryptBlock(blockNumber)
		decrypted = append(decrypted, decryptedBlock...)
	}

	return decrypted
}

// DecryptBlock decrypts the given block.
// It truncates the message at the block after this one to
// use the padding oracle attack.
func DecryptBlock(blockNumber int) []byte {
	decrypted := make([]byte, BlockSize)
	// Decrypt bytes one by one.
	for i := 0; i < BlockSize; i++ {
		cipher := cryptoutil.NewCipher(EncryptedCipherText)[:BlockSize*(blockNumber+2)]

		// Insert previously found bytes.
		for j := 0; j < i; j++ {
			cipher[len(cipher)-BlockSize-1-j] = cipher[len(cipher)-BlockSize-1-j] ^ decrypted[BlockSize-1-j] ^ byte(i+1)
		}

	guessLoop:
		// Guess the current byte.
		for g := byte(0); g <= 255; g++ {
			encryptedGuess := make([]byte, len(cipher))
			copy(encryptedGuess, cipher)
			encryptedGuess[len(encryptedGuess)-BlockSize-1-i] = encryptedGuess[len(encryptedGuess)-BlockSize-1-i] ^ g ^ byte(i+1)
			statusCode := SendRequest(encryptedGuess)
			if statusCode == InvalidMessageErrorCode {
				fmt.Printf("[%d] %d\n", i, g)
				decrypted[BlockSize-1-i] = g
				break guessLoop
			} else if statusCode == 200 {
				fmt.Printf("[%d] Got HTTP200 for %d\n", i, g)
				if g > 1 {
					decrypted[BlockSize-1-i] = g
					break guessLoop
				}
			}
		}
	}

	return decrypted
}

// FindPaddingLength sends up to 256 guesses for the last byte of the
// decrypted message.
// That last byte will be the pad length.
func FindPaddingLength() byte {
	for i := byte(0); i <= 255; i++ {
		encryptedGuess := cryptoutil.NewCipher(EncryptedCipherText)
		encryptedGuess[3*BlockSize-1] = encryptedGuess[3*BlockSize-1] ^ 1 ^ i
		statusCode := SendRequest(encryptedGuess)
		if statusCode == InvalidMessageErrorCode {
			return i
		}
	}

	return 0
}

// SendRequest tries the given cipher text and returns the status code.
func SendRequest(cta []byte) int {
	r, err := http.Get(SiteURL + hex.EncodeToString(cta))
	if err != nil {
		fmt.Printf("Err: %s\n", err.Error())
	}

	// fmt.Printf("%x: %d\n", cta, r.StatusCode)
	return r.StatusCode
}
