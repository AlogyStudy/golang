package main

import (
	"fmt"
)

func main () {

	var a int // 变量声明
	a = 10 // 变量赋值

	var b int = 20 // 变量声明的同时赋值
	var c = 1 // 变量声明与赋值，由系统推荐是那种类型

	d := 10 // 函数中的变量声明与赋值的最简写法

	fmt.Println(a, b, c, d)

}