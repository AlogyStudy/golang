package main

import (
	"fmt"
)

// 第一种形式
// func main () {
// 	a := 1
// 	for {
// 		a++
// 		if a > 3 {
// 			break
// 		}
// 		fmt.Println(a) // 2, 3
// 	}
// 	fmt.Println(a) // 4
// }


// 第二种形式
// func main () {
// 	a := 1
// 	for a <= 3 {
// 		a++
// 		fmt.Println(a) // 2, 3, 4
// 	}
// 	fmt.Println(a) // 4
// }

// 第三种形式
func main () {
	a := "string"
	len := len(a)
	for i := 0; i < len; i++ {
		// a++
		fmt.Println(i) // 2, 3, 4
	}
	// fmt.Println(a) // 4
}
