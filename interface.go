package main

import (
	"fmt"
)

// Go语言中所有类型都实现空接口
// type empty interface {
// }

type USB interface {
	Name() string
	// Connect()
	Connecter
}

type Connecter interface {
	Connect()
}

type Phone struct {
	name string
}

func (pc Phone) Name() string {
	return pc.name
}

type TVConnector struct {
	name string
}

func (pc Phone) Connect() {
	fmt.Println("Connect: ", pc.name)
}

func main () {
	// var a USB
	pc := Phone{"phone"}
	var a Connecter
	// a = Phone{name: "phone"}
	a = Connecter(pc)
	a.Connect()
	Disconnect(a)

	// 空接口的判断
	var b interface{}
	fmt.Println(b == nil) // true

	var p *int = nil
	b = p
	fmt.Println(b == nil) // false
}

func Disconnect(usb interface{}) { // interface{} 空接口
	// if pc,ok := usb.(Phone); ok { // 类型判断
	// 	fmt.Println("Disconnect：", pc.name)	
	// 	return
	// }
	switch v := usb.(type) {
		case Phone:
			fmt.Println("Disconnect：", v.name)
		default :
			fmt.Println("Unknown")
	}
	fmt.Println("Disconnect")
}
