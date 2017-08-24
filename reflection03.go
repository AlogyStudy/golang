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

func main () {
	u := User{1, "Ok", 18}
	Set(&u)
	fmt.Println(u)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if reflect.Ptr == v.Kind() && !v.Elem().CanSet() { // 指针是否正确
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem() // 重写赋值
	}

	f := v.FieldByName("Name") // 获取字段
	if !f.IsValid() { // 判读是否取到当前字段
		fmt.Println("BAD")
		return
	}
	if f.Kind() == reflect.String { // 类型判断
		f.SetString("MM") // 重新赋值
	}

}
