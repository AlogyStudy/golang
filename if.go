package main

import (
	"fmt"
)

func main () {
	a := 10
	if a := 0; a > 0 {
		fmt.Println(a)
	} else if a == 0 {
		fmt.Println(0111) // 73
	}
	fmt.Println(a) // 10
}
