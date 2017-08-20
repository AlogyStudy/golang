package main

import (
	"fmt"
	// "sort"
)

func main () {
	// var m map[int]string
	// m = map[int]string{}
	// m = make(map[int]string)
	// var m1 map[int]string = make(map[int]string)


	// m2 := make(map[int]string)
	// m2[1] = "OK"
	// delete(m2, 1)
	// a := m2[1]

	// fmt.Println(a) // map[]


	// var m map[int]map[int]string
	// m = make(map[int]map[int]string)
	// // m[1] = make(map[int]string)
	// // m[1][1] = "OK"
	// a, ok := m[2][1]
	// if !ok {
	// 	m[2] = make(map[int]string)
	// }
	// m[2][1] = "X"
	// a = m[2][1]
	// fmt.Println(a, ok)


	// for i,v := range  slice {
	// 	// slice[i]
	// }

	// for k,v := range map {
	// 	// map[k]
	// }
	

	// sm := make([]map[int]string, 5)

	// for i := range sm {
	// 	sm[i] = make(map[int]string, 1)
	// 	sm[i][1] = "ok"
	// 	fmt.Println(sm[i])
	// }
	// fmt.Println(sm)
	

	// m := map[int]string{1: "A", 2: "B", 3: "C", 4: "D"}
	// s := make([]int, len(m))
	// i := 0
	// for k,_ := range m {
	// 	s[i] = k
	// 	i++
	// }
	// sort.Ints(s)
	// fmt.Println(s)

	m1 := map[int]string{1: "A", 2: "B", 3: "C"}
	m2 := make(map[string]int)
	// m2 := map[string]int{"A": 1, "B": 2, "C": 3}

	for k,v := range m1 {
		// fmt.Println(k, v)
		m2[v] = k
	}
	fmt.Println(m1)
	fmt.Println(m2)

}
