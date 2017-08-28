package main

import (
	"fmt"
)

// 创建一个goroutine，与主线程按顺序相互发送信息若干次并打印
var c chan string

func Pingpong() {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("Pingpong: #%d", i)
		i++
	}
}

func main() {
	c = make(chan string)
	go Pingpong()

	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("Main: #%d", i)
		fmt.Println(<-c)
	}
}
