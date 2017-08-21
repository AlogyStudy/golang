package main

import (
	"fmt"
)


type person struct {
	name string
	age int
	contact struct {
		phone,city string
	}
}

type p1 struct {
	string
	int
}

func main () {
	b := person {
		name: "yellow",
		age: 18,
	}
	b.contact.phone = "123123"
	b.contact.city = "xiamen"

	a := &struct {
		name string
		age int
	} {
		name: "tan",
		age: 19,
	}

	c := p1{"cyan", 20}
	var d p1
	d = c

	fmt.Println(a, b, c, d)
}
