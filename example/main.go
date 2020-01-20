package main

import (
	"flag"
	"fmt"
	"github.com/chaseisabelle/flagz"
)

func main() {
	var arr flagz.Flagz

	flag.Var(&arr, "foo", "bla")

	flag.Parse()

	fmt.Printf("%+v\n%+v\n", arr.Array(), arr.String())
}
