package main

import (
	"fmt"
)

// const a int = 1
// const b string  = 'A'

// const (
// 	c = a
// 	d = a + 2
// 	e = 3
// )

// const d, f, g, h = 10, 11, 12, 13

// const (
// 	x, y, z = 'x', 'y', 'z'
// )

// -----

// const (
// 	a = 1
// 	b
// 	c
// )

// func main () {
// 	fmt.Println(a, b, c) // 1, 1, 1
// }

// -----

// var ss = "123"

// const (
// 	a, b = 1, "xixi"
// 	c, d
// )

// func main () {
// 	fmt.Println(a, b, c, d)
// }



const (
	_A = "A"
	_B = iota
	_C = "B"
	_D = iota
)
const (
	_E = iota
)
func main () {
	fmt.Println(_A, _B, _C, _D, _E) // A 1 B 3 0
}

