package main 

import (
	"fmt"
)

var (
	// 使用常规方式
	aaa = "hello"
	// 使用并行方式以及类型推断
	aa, b = 1, 2
	// cc := 2 // 不可以省略 var
)


func main () {
	// 多个变量的声明
	var a, b, c, d int
	// 多个变量的赋值
	a, b, _, d = 1, 2, 3, 4

	// 多个变量声明的同时赋值
	var e, f, g, h int = 5, 6, 7, 8
	// 省略变量类型，由系统推断
	var i, j, k, l = 9, 10, 11, 12
	// 多个变量声明与赋值的最简写法
	i, m, n, o := 13, 14, 15, 16

  // _ := 10 // 空白符号，省略该表达式赋值(应用函数返回值)

  fmt.Println(aaa, aa, a, c, b, d, e, f, g, h, i, m, n, o, j, k, l )
}