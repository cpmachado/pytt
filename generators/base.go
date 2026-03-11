package generators

import (
	"iter"
)

type Pytt struct {
	A, B, C int
}

func Base(n int, k int) iter.Seq2[int, Pytt] {
	return func(yield func(int, Pytt) bool) {
		i := 0
		for a := range n {
			for b := range a {
				for c := range b {
					if i >= k {
						return
					}
					if a*a == b*b+c*c {
						if !yield(i, Pytt{A: a, B: b, C: c}) {
							return
						}
						i++
					}
				}
			}
		}
	}
}
