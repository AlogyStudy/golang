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

func (u User) Hello(name string) {
	fmt.Println("Hello", name, ", my name is", u.Name)
}

func main () {
	u := User{1, "OK", 123}
	v := reflect.ValueOf(u)

	mv := v.MethodByName("Hello") // 获取函数名
	args := []reflect.Value{reflect.ValueOf("alogy")} // 传递参数

	mv.Call(args)

	u.Hello("alogy")
}
