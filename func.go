package main

import (
	"fmt"
)

func main () {
	// test(10)
	
	a := func () {
		fmt.Println("匿名函数")
	}
	a()

	var t = closure(10)(100)
	fmt.Println(t)
}


// func test (a int, b string) (int, string, int) {

// }

// func test (a ...int) (a, b, c int) {
// 	return a, b, c
// }

// func test (a int) {
// 	fmt.Println(a)
// }
// 

func closure (x int) func (int) int {
	fmt.Println(&x)
	return func (y int) int {
		fmt.Println(&x)
		return x + y
	}
}
