package main

import (
	"fmt"

	"github.com/flambra/helpers/test"
)

func main() {
	test.Print()

	oct := test.Octopus{
		Name:  "Jesse",
		Color: "orange",
	}

	fmt.Println(oct.String())

	oct.Reset()

	fmt.Println(oct.String())
}
