package main

import (
	"fmt"
)

func main() {
	// c := make(chan bool, 1) // 有缓存 异步
	c := make(chan bool) // 无缓存的时候是阻塞的
	go func () {
		fmt.Println("Go...")	
		<- c // 读取
	}()
	c <- true // 传进去
}
