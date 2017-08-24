package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

type Manager struct {
	User
	title string
}

func main () {
	// 反射会将匿名字段当作独立字段来处理
	m := Manager{User: User{1, "OK", 12}, title: "123123"}
	t := reflect.TypeOf(m)
	fmt.Println(t.Field(0))

	// 取匿名当中的字段
	fmt.Println(t.FieldByIndex([]int{0, 0}))
}
