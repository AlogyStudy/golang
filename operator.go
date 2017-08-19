package main

import (
	"fmt"
)

func main () {
	fmt.Println(1 ^ 2) // 二元运算符
	fmt.Println(^2) // 一元运算符
	fmt.Println(6&11) // 2
	fmt.Println(6|11) // 15
	fmt.Println(6^11) // 13 
	fmt.Println(6&^11) // 4
	a := 0
	if a > 0 && (10 / a ) > 1 {
		fmt.Println("Ok")
	}
}

/*

	6: 0110
 11: 1101
------------
  &	 0010  // 2
  |  1111  // 15
  ^  1101  // 13
  &^ 0100  // 4

	6 -> 110
	5 -> 101
	4 -> 100

13 / 2 = 1 // 6
1101
 */
