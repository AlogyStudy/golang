package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go...")
		c <- true // 存 // 声明的时候是bool
	}()
	<-c // 取 // 消息存取，阻塞执行
}
