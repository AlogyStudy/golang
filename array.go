package main

import (
	"fmt"
)

func main () {
	// var a [2]string
	// var b [1]int
	// c := [2]int{11, 12}
	// d := [20]int{19: 1}
	// e := [...]int{1, 2, 3, 4, 5}
	// f := [...]int{0: 11, 1: 22, 2: 33}
	// b[0] = 10
	// a[1] = "100"

	// fmt.Println(a, b, c, d, e, f)
	// 
	// a := [...]int{99: 1}

	// var p *[100]int = &a
	// x, y := 10, 20

	// var pp [...]*int = {&x, &y}
	// fmt.Println(pp)
	// 
	// a := [...]int{1, 2}
	// b := [...]int{1, 3}
	// fmt.Println(a == b)
	// 
	// p := new([10]int)
	// p[1] = 2
	// fmt.Println(*p)
	// 
	// 多维数组
	a := [2][3]int{
		{1, 1, 1},
		{2, 2, 2},
	}
	fmt.Println(a)
}
