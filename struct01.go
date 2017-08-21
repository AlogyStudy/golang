package main

import (
	"fmt"
)

type human struct {
	Sex int
}

type teacher struct {
	human
	name string
	age int
}

type student struct {
	human
	name string
	age int
}

func main () {
	// a := teacher{name: "cyan", age: 20, human{sex: 0}}
	a := teacher{name: "cyan", age: 20, human: human{Sex: 0}}
	b := student{name: "pink", age: 22, human: human{Sex: 1}}

	a.name = "xixi"
	a.age = 23
	// a.Sex = 100
	a.human.Sex = 200

	fmt.Println(a, b)
}
