package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for v := range c {
			fmt.Println(v)
		}	
	}()

	for i := 0; i < 10; i++ {
		select {
			case c <- 0:
			case c <- 1:	
		}
	}

	select{} // 阻塞main函数退出，卡死main函数，场景经常使用在事件循环
}
