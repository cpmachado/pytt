package generators

import (
	"iter"
	"math/big"
)

var one, two *big.Int = big.NewInt(1), big.NewInt(2)

func Euclid(k int) iter.Seq2[int, Pytt] {
	return func(yield func(int, Pytt) bool) {
		i := 0
		one := big.NewInt(1)
		for m := big.NewInt(2); ; m.Add(m, one) {
			for n := big.NewInt(1); n.Cmp(m) < 0; n.Add(n, one) {
				if i >= k {
					return
				}
				if checkValidMN(m, n) {
					p := generateTriple(m, n)

					if !yield(i, p) {
						return
					}
					i++
				}
			}
		}
	}
}

func generateTriple(m, n *big.Int) Pytt {
	m2 := new(big.Int).Mul(m, m)
	n2 := new(big.Int).Mul(n, n)
	a := new(big.Int).Sub(m2, n2)
	b := new(big.Int).Mul(two, new(big.Int).Mul(m, n))
	c := new(big.Int).Add(m2, n2)

	if a.Cmp(b) > 0 {
		a, b = b, a
	}

	return Pytt{A: *a, B: *b, C: *c}
}

func checkValidMN(m, n *big.Int) bool {
	return new(big.Int).Rem(new(big.Int).Sub(m, n), two).Cmp(one) == 0 && new(big.Int).GCD(nil, nil, m, n).Cmp(one) == 0
}
