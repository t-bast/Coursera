// Package week4 contains the implementation of
// a padding oracle attack.
package week4

const (
	// SiteURL is the hostname of the vulnerable website.
	SiteURL = "http://crypto-class.appspot.com/po?er="

	// InvalidPadErrorCode is the http status returned in case of
	// invalid padding.
	InvalidPadErrorCode = 403

	// InvalidMessageErrorCode is the http status returned in case
	// the padding is correct but the message is invalid.
	InvalidMessageErrorCode = 404

	// BlockSize is 16 because we know that AES CBC is used.
	BlockSize = 16

	// EncryptedCipherText is the cipher text we want to decrypt
	// thanks to the padding oracle vulnerability.
	EncryptedCipherText = "f20bdba6ff29eed7b046d1df9fb7000058b1ffb4210a580f748b4ac714c001bd4a61044426fb515dad3f21f18aa577c0bdf302936266926ff37dbf7035d5eeb4"
)
