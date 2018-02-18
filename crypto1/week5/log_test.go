package week5

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

// For our tests, 12 = 4^9 [13] and we want to find x=3 (because 4^3=4^9 [13]).
// 1 <= x <= 2^4 so B=2^2.
var (
	p = big.NewInt(13)
	g = big.NewInt(4)
	h = big.NewInt(12)
)

func TestComputeHashTable(t *testing.T) {
	x1s := computeHashTable(h, g, p, 2)

	assert.Len(t, x1s, 5)
	assert.Equal(t, 0, x1s["12"])
	assert.Equal(t, 1, x1s["3"])
	assert.Equal(t, 2, x1s["4"])
	assert.Equal(t, 3, x1s["1"])
	assert.Equal(t, 4, x1s["10"])
}

func TestComputeLog(t *testing.T) {
	x, _ := ComputeLog(p.String(), g.String(), h.String(), 2)
	assert.Equal(t, uint64(3), x)
}
