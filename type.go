package main

import (
	"fmt"
)

func main () {
	var a float32 = 100.01
	fmt.Println(a) // 100.01
	b := int(a)
	fmt.Println(b) // 100
	var c int8 = 5
	d := string(c)
	fmt.Println(d + "---")
}
