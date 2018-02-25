package week6

import (
	"fmt"
	"math/big"
)

// Factor factorizes an RSA modulus when the primes are too close.
// We assume that |p-q| < 2 * N^(1/4).
func Factor(n *big.Int) (*big.Int, *big.Int) {
	a := calcA(n)
	x := calcX(a, n)
	return calcFactors(a, x)
}

// Factor2 factorizes an RSA modulus when the primes are close.
// We assume that |p-q| < 2^11 * N^(1/4).
func Factor2(n *big.Int) (*big.Int, *big.Int) {
	nn := new(big.Int).Sqrt(n)
	a := new(big.Int).Add(nn, big.NewInt(1))
	for {
		x := calcX(a, n)
		p, q := calcFactors(a, x)
		pq := new(big.Int).Mul(p, q)
		if pq.Cmp(n) == 0 {
			return p, q
		}

		a = a.Add(a, big.NewInt(1))
	}
}

// Factor3 factorizes an RSA modulus when the primes are close.
// We assume that |3p-2q| < N^(1/4).
func Factor3(n *big.Int) (*big.Int, *big.Int) {
	n24 := new(big.Int).Mul(big.NewInt(24), n)
	n24sqrt := new(big.Int).Sqrt(n24)

	// a = 3p+2q
	// We can't use 3p+2q/2 because it's not an integer.
	a := new(big.Int).Add(n24sqrt, big.NewInt(1))
	a2 := new(big.Int).Mul(a, a)

	// We define x such that:
	//	* 6p=a-x
	//	* 4q=a+x
	// We find that x=sqrt(a^2-24N).
	x := new(big.Int).Sqrt(new(big.Int).Sub(a2, n24))

	p := new(big.Int).Div(new(big.Int).Sub(a, x), big.NewInt(6))
	q := new(big.Int).Div(new(big.Int).Add(a, x), big.NewInt(4))

	return p, q
}

// Decrypt decrypts an RSA-encrypted message once we have the
// prime factors.
func Decrypt(p, q, e, c *big.Int) (string, error) {
	n := new(big.Int).Mul(p, q)

	phi := new(big.Int).Mul(
		new(big.Int).Sub(p, big.NewInt(1)),
		new(big.Int).Sub(q, big.NewInt(1)),
	)

	d := new(big.Int).ModInverse(e, phi)

	decrypted := new(big.Int).Exp(c, d, n).Bytes()
	if decrypted[0] != 0x02 {
		return "", fmt.Errorf("Invalid preamble: %x", decrypted[0])
	}

	msgIndex := 0
	for msgIndex = 0; msgIndex < len(decrypted); msgIndex++ {
		if decrypted[msgIndex] == 0x00 {
			msgIndex++
			break
		}
	}

	msg := decrypted[msgIndex:]
	return fmt.Sprintf("%s", msg), nil
}

// calcA returns ceil(sqrt(n)).
func calcA(n *big.Int) *big.Int {
	almostA := new(big.Int).Sqrt(n)
	a := new(big.Int).Add(almostA, big.NewInt(1))
	return a
}

// calcX returns sqrt(a^2-n).
func calcX(a, n *big.Int) *big.Int {
	a2 := new(big.Int).Mul(a, a)
	a2n := new(big.Int).Sub(a2, n)
	return new(big.Int).Sqrt(a2n)
}

// calcFactors calculates the factors from a and x.
// Depending on how close p and q are, you might need to
// verify the factorization.
func calcFactors(a, x *big.Int) (*big.Int, *big.Int) {
	p := new(big.Int).Add(a, x)
	q := new(big.Int).Sub(a, x)
	return p, q
}
