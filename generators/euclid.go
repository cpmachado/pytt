package generators

import "iter"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Euclid(k int) iter.Seq2[int, Pytt] {
	return func(yield func(int, Pytt) bool) {
		i := 0
		for m := 2; ; m++ {
			for n := 1; n < m; n++ {
				if i >= k {
					return
				}
				if (m-n)%2 == 1 && gcd(m, n) == 1 {
					a := m*m - n*n
					b := 2 * m * n
					c := m*m + n*n

					if a > b {
						a, b = b, a
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
