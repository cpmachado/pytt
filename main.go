package main

import (
	"flag"
	"fmt"
)

func main() {
	n := 100
	h := false

	flag.IntVar(&n, "n", n, "upper bound")
	flag.BoolVar(&h, "t", h, "print header")
	flag.Parse()

	if h {
		fmt.Println("a,b,c")
	}

	for a := range n {
		for b := range a {
			for c := range b {
				if a*a == b*b+c*c {
					fmt.Printf("%d,%d,%d\n", a, b, c)
				}
			}
		}
	}
}
