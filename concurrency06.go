package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)
	go func() {
		for { // 通过死循环来不断发送和接收chan
			select {
			case v, ok := <-c1 :
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <- c2 :
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()
		
	c1 <- 1
	c2 <- "hello"
	c1 <- 2
	c2 <- "zf"

	close(c1)
	close(c2)

	for i := 0; i < 2; i++ {
		<-o
	}

}
