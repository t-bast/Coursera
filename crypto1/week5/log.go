package week5

import (
	"fmt"
	"math/big"

	"github.com/pkg/errors"
)

// ComputeLog computes x where h = g^x [p]
func ComputeLog(p, g, h string, bpow uint) (uint64, error) {
	np, ok := new(big.Int).SetString(p, 10)
	if !ok {
		return 0, errors.New("Invalid number")
	}

	ng, ok := new(big.Int).SetString(g, 10)
	if !ok {
		return 0, errors.New("Invalid number")
	}

	nh, ok := new(big.Int).SetString(h, 10)
	if !ok {
		return 0, errors.New("Invalid number")
	}

	x1s := computeHashTable(nh, ng, np, bpow)

	b := big.NewInt(1 << bpow)
	gb := new(big.Int).Exp(ng, b, np)
	x0, x1 := findMatch(gb, np, bpow, x1s)

	res := uint64(x0)*(1<<bpow) + uint64(x1)

	return res, nil
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

func findMatch(gb, p *big.Int, bpow uint, xMap map[string]int) (int, int) {
	gbpow := big.NewInt(1)

	for i := 0; i <= (1 << bpow); i++ {
		if match, ok := xMap[gbpow.String()]; ok {
			return i, match
		}

		gbpow = gbpow.Mul(gbpow, gb)
	}

	fmt.Println("Could not find a solution")
	return 0, 0
}
