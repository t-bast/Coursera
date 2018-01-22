// Programming assignments for coursera crypto courses.
package main

import (
	"fmt"

	"github.com/t-bast/coursera/crypto1/week1"
)

func main() {
	var message week1.Message = "attack"
	b, _ := message.Bytes()
	printBinary(b)

	var cipher week1.Cipher = "315c4eea"
	b, _ = cipher.Bytes()
	printBinary(b)
}

func printBinary(b []byte) {
	for _, n := range b {
		fmt.Printf("%08b", n)
	}
	fmt.Println()
}
