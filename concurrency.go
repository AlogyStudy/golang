package main

import (
	"fmt"
	"time"
)

func main () {
	fmt.Println(2 * time.Second) // 2s
	go Go()
	time.Sleep(2 * time.Second) // 延迟2s
}

func Go () {
	fmt.Println("Go...")
}