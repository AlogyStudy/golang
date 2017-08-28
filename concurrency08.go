package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	select {
		case v := <-c:
			fmt.Println(v)
		case t := <-time.After(3 * time.Second):
			fmt.Println(t)
			fmt.Println("Timeout")
	}
}
