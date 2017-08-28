package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //  使用多核分配的时候，任务分配时不一定的
	wg := sync.WaitGroup{} // 内置包sync使用
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}

	wg.Wait()
}

func Go(wg *sync.WaitGroup, idx int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(idx, a)

	wg.Done()
}
