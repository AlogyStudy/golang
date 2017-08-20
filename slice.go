package main

import (
	"fmt"
)

func main () {

	// 声明方法1：
	// var s1 []int // 中括号中没有数字或`...`
	// fmt.Println(s1) // []

	// a := [10]int{}
	// fmt.Println(a)
	// s1 := a[5: len(a)] // 包含起始索引，不包含终止索引 // a[5 6 7 8 9]
	// s2 := a[5: ] // 包含起始索引，不包含终止索引 // a[5 6 7 8 9]
	// fmt.Println(s1, s2)


	// s1 := make([]int, 3, 10) // 10小块连续的内存,如果slice超过10，内存卡会继续申请，重新生成内存地址
	// s2 := make([]int, 10) // cap不给定，是slice的最大长度
	// fmt.Println(len(s1), cap(s1), s1) // 3 10 [0 0 0]
	// fmt.Println(len(s2), cap(s2), s2) // 10 10 [0 0 0 0 0 0 0 0 0 0]


	// a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	// sa := a[2: 5]
	// sb := a[3: 5]
	// fmt.Println(string(sa), string(sb))


	// a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	// sa := a[2: 5]
	// sb := a[3:5]
	// sb_1 := sa[1: 3]
	// sb_2 := sa[3: 5]
	// fmt.Println(string(sa), string(sb), string(sb_1), string(sb_2))// cde de de fg

	// s1 := make([]int, 3, 6)
	// fmt.Println("%p\n", s1)
	// s1 = append(s1, 1, 2, 3, 4)

	// fmt.Println("%v %p\n", s1)
	
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	s2 := []int{8, 9}
	// copy(s1, s2)
	copy(s1, s2[1: 2])
	fmt.Println(s1, s2)

	s11 := []int{1, 2, 3, 4}
	s22 := s11[:]
	fmt.Println(s11, s22)

}