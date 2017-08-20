package main 

import (
	"fmt"
)

// func main () {
// 	a := 1
// 	switch a {
// 	case 0:
// 		fmt.Println("a=0")
// 	case 1:
// 		fmt.Println("a=1")	
// 	}
// 	fmt.Println(a)
// }


// func main () {
// 	a := 1
// 	switch {
// 	case a >= 0:
// 		fmt.Println("a>=0")
// 		fallthrough
// 	case a >= 1:
// 		fmt.Println("a>=1")
// 	}
// 	fmt.Println(a)
// }



func main () {
	switch a := 1; {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough
	case a >= 1:
		fmt.Println("a>=1")
	default:
		fmt.Println("none")	
	}
	// fmt.Println(a) // undefined: a
}
