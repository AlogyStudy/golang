package main

import (
	"fmt"
)

func main () {
	// fmt.Println("a")
	// defer fmt.Println("b")
	// defer fmt.Println("c")
	// a, c, b
	
	
	// for i := 0; i < 3; i++ {
	// 	defer fmt.Println(i) // 2 1 0
	// }
	
	// for i := 0; i < 3; i++ {
	// 	defer func () {
	// 		fmt.Println(i) // 3 3 3
	// 	}()
	// }
	
	A()
	B()
	C()

}

func A () {
	fmt.Println("func A")
}

func B () {
	defer func () {
		if err := recover(); err != nil {
			fmt.Println("Recover")
		}
	}()
	panic("Panic in B")
}

func C () {
	fmt.Println("func C")
}
