package main

import (
	"flag"
	"fmt"
)

const usageContent = `%s is a simple tool to compute pythagorean triplets, in csv form.

Usage of %s:
  -n int
    	upper bound (default 100)
  -t	print header
`

func main() {
	n := 100
	h := false

	flag.IntVar(&n, "n", n, "upper bound")
	flag.BoolVar(&h, "t", h, "print header")
	flag.Usage = func() {
		fmt.Printf(usageContent, flag.CommandLine.Name(), flag.CommandLine.Name())
	}

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
