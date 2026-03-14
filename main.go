package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/cpmachado/pytt/generators"
)

func main() {
	n := 100
	h := false
	v := false

	flag.IntVar(&n, "n", n, "number of triplets")
	flag.BoolVar(&h, "t", h, "print header")
	flag.BoolVar(&v, "v", v, "display version")

	flag.Parse()

	if v {
		version()
		os.Exit(0)
	}

	if h {
		fmt.Println("a,b,hypotenuse")
	}

	gen := generators.Euclid

	for _, v := range gen(n) {
		a, b, c := v.A, v.B, v.C
		fmt.Printf("%s,%s,%s\n", a.String(), b.String(), c.String())
	}
}

func version() {
	vsn := "dev"
	if info, ok := debug.ReadBuildInfo(); ok {
		vsn = info.Main.Version
	}
	fmt.Printf("pytt-%s\n", vsn)
}
