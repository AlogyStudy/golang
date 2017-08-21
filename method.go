package main

import (
	"fmt"
)

type Test struct {
	name string
}

type Person struct {
	name string
}

func main () {
	t := Test{}
	t.Print()
	fmt.Println(t.name)

	p := Person{}
	p.Print()
	fmt.Println(p.name)
}

func (t *Test) Print() {
	t.name = "red"
	fmt.Println("Test")
}

func (p Person) Print() {
	fmt.Println("Person")
}
