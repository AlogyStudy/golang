package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go...")
		c <- true
		close(c) // 关闭chan // 没有明确关闭，会出现死锁，崩溃退出
	}()
	// <- c
	for v := range c {
		fmt.Println(v)
	}
}
