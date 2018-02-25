package week6

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcA(t *testing.T) {
	n := big.NewInt(35)
	a := calcA(n)
	assert.Equal(t, int64(6), a.Int64())
}

func TestCalcX(t *testing.T) {
	n := big.NewInt(35)
	a := calcA(n)
	x := calcX(a, n)
	assert.Equal(t, int64(1), x.Int64())
}

func TestFactor(t *testing.T) {
	n := big.NewInt(35)
	p, q := Factor(n)
	assert.Equal(t, int64(7), p.Int64())
	assert.Equal(t, int64(5), q.Int64())
}

func TestFactor3(t *testing.T) {
	n := big.NewInt(2039652913367)
	p, q := Factor3(n)
	assert.Equal(t, int64(1166083), p.Int64())
	assert.Equal(t, int64(1749149), q.Int64())
}
