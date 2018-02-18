package week5

import (
	"fmt"
	"math/big"

	"github.com/pkg/errors"
)

// ComputeLog computes x where h = g^x [p]
func ComputeLog(p, g, h string, bpow uint) (string, error) {
	np, ok := new(big.Int).SetString(p, 10)
	if !ok {
		return "", errors.New("Invalid number")
	}

	ng, ok := new(big.Int).SetString(g, 10)
	if !ok {
		return "", errors.New("Invalid number")
	}

	nh, ok := new(big.Int).SetString(h, 10)
	if !ok {
		return "", errors.New("Invalid number")
	}

	x1s := computeHashTable(nh, ng, np, bpow)
	x0 := findMatch(ng, np, x1s)
	fmt.Println(x0.String())

	return "", nil
}

func computeHashTable(h, g, p *big.Int, bpow uint) map[string]int {
	res := make(map[string]int)
	gpow := big.NewInt(1)

	// For i=0..2^bpow, we compute h/(g^i) [p]
	// and store the result.
	for i := 0; i <= (1 << bpow); i++ {
		// Compute the inverse of g^i [p].
		ginv := new(big.Int).ModInverse(gpow, p)

		// Multiply it by h modulo p.
		hres := new(big.Int).Mul(h, ginv)
		hres = hres.Mod(hres, p)

		// Store the result.
		res[hres.String()] = i

		// Prepare next iteration.
		gpow = gpow.Mul(gpow, g)
	}

	return res
}

func findMatch(g, p *big.Int, xMap map[string]int) *big.Int {
	return nil
}
