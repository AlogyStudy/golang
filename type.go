package main

import (
	"fmt"
	"strconv"
)

func main () {
	var a float32 = 100.01
	fmt.Println(a) // 100.01
	b := int(a)
	fmt.Println(b) // 100
	var c int = 3
	// d := string(c)
	d := strconv.Itoa(c)
	c, _ = strconv.Atoi(d)
	fmt.Println(d + "---")
	fmt.Println(c)
}
