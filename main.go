// Programming assignments for coursera cryptography courses.
package main

import (
	"fmt"

	"github.com/t-bast/coursera/crypto1/cryptoutil"
	"github.com/t-bast/coursera/crypto1/week2/ctr"
)

func main() {
	b, err := ctr.Decrypt(
		cryptoutil.NewCipher("hex key here"),
		cryptoutil.NewCipher("hex cipher text here"))

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Println(cryptoutil.Cipher(b).ASCII())
}
