package generators

import "iter"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Euclid(limit int, k int) iter.Seq2[int, Pytt] {
	return func(yield func(int, Pytt) bool) {
		i := 0
		for m := 2; m <= limit; m++ {
			for n := range m {
				if i >= k {
					return
				}
				if gcd(m, n) == 1 && (m-n)%2 == 1 {
					a := m*m - n*n
					b := 2 * m * n
					c := m*m + n*n
					if c > limit || a > limit || b > limit {
						return
					}

					if !yield(i, Pytt{a, b, c}) {
						return
					}
					i++
				}
			}
		}
	}
}
