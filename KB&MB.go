package main

import (
	"fmt"
)

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
)

func main () {
	fmt.Println(B, KB, MB, GB, TB)
}

// 0001

// 001000000000 // 512
