package main

import (
	"flag"
	"fmt"

	"github.com/vestiaire-collective/flagz"
)

func main() {
	var stringz flagz.Flagz
	var boolz flagz.Flagz
	var intz flagz.Flagz
	var floatz flagz.Flagz

	flag.Var(&stringz, "string", "strings")
	flag.Var(&boolz, "bool", "bools")
	flag.Var(&intz, "int", "ints")
	flag.Var(&floatz, "float", "floats")

	flag.Parse()

	strings := stringz.Stringz()
	bools, boolErr := boolz.Boolz()
	ints, intErr := intz.Intz()
	floats, floatErr := floatz.Floatz()

	fmt.Printf("%+v\n", strings)
	fmt.Printf("%+v\n%+v\n", bools, boolErr)
	fmt.Printf("%+v\n%+v\n", ints, intErr)
	fmt.Printf("%+v\n%+v\n", floats, floatErr)
}
