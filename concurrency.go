package main

import (
	"fmt"
	"time"
)

func main () {
	fmt.Println(2 * time.Second) // 2s
	go Go()
	time.Sleep(2 * time.Second) // 延迟2s
	// 在main函数运行Sleep的时候，Go函数也运行了,执行完之后，退出。
}

func Go () {
	fmt.Println("Go...")
}
