package main

import (
	"fmt"
)

func main () {
	a := 1
	a--
	var p *int = &a
	fmt.Println(p) // 0xc0420361d0
	fmt.Println(*p) // 1
}