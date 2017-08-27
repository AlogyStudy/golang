package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //  使用多核分配的时候，任务分配时不一定的
	c := make(chan bool, 10) // 设置缓存
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}

	for i := 0; i < 10; i++ {
		<-c
	}
}

func Go(c chan bool, idx int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(idx, a)

	c <- true
}
