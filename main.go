package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"runtime/debug"

	"github.com/cpmachado/pytt/generators"
)

func main() {
	n := math.MaxInt
	k := 100
	h := false
	v := false

	flag.IntVar(&n, "n", n, "upper bound")
	flag.IntVar(&k, "k", k, "number of triplets")
	flag.BoolVar(&h, "t", h, "print header")
	flag.BoolVar(&v, "v", v, "display version")

	flag.Parse()

	if v {
		vsn := "dev"
		if info, ok := debug.ReadBuildInfo(); ok {
			vsn = info.Main.Version
		}
		fmt.Printf("pytt-%s\n", vsn)
		os.Exit(0)
	}

	if h {
		fmt.Println("a,b,c")
	}

	gen := generators.Euclid

	for _, v := range gen(n, k) {
		a, b, c := v.A, v.B, v.C
		fmt.Printf("%d,%d,%d\n", a, b, c)
	}
}
