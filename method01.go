package main

import (
	"fmt"
)

// 类型别名不会拥有底层类型所附带的方法
// type TZ int

// 属性的访问范围是在`package`中的可以访问的，如果需要在外部包中访问，需要大写字母
type A struct {
	name string
}

func main () {
	a := A{}
	a.Print()
	fmt.Println(a.name)
}

func (a *A)  Print() {
	a.name = "123"
	fmt.Println(a.name)
}
