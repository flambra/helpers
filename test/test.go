package test

import "fmt"

type Octopus struct {
	Name  string
	Color string
}

func (o Octopus) String() string {
	return fmt.Sprintf("The octopus's name is %q and is the color %s.", o.Name, o.Color)
}

func (o *Octopus) Reset() {
	o.Name = ""
	o.Color = ""
}

func Print() {
	fmt.Println("Helpers module is working!")
}