// Programming assignments for coursera cryptography courses.
package main

import (
	"fmt"

	"github.com/t-bast/coursera/crypto1/week4"
)

func main() {
	decryptedBytes := week4.Decrypt()
	fmt.Println("Decrypted message:")
	fmt.Printf("%s\n", decryptedBytes)
}
