package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
)

const usageContent = `%s is a simple tool to compute pythagorean triplets, in csv form.

Usage of %s:
  -n int
    	upper bound of each element (default 100)
  -t	print header
  -v	print version
`

func main() {
	n := 100
	h := false
	v := false

	flag.IntVar(&n, "n", n, "upper bound of each element")
	flag.BoolVar(&h, "t", h, "print header")
	flag.BoolVar(&v, "v", v, "display version")
	flag.Usage = func() {
		fmt.Printf(usageContent, flag.CommandLine.Name(), flag.CommandLine.Name())
	}

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
