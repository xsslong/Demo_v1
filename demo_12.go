package main

import "fmt"

func main() {
	a := []int{3, 4, 6, 7, 5, 9, 1}
	s1 := a[2:5] // [6  7  5]
	s2 := a[4:6] //		  [5  9]
	fmt.Println(s1)
	fmt.Println(s2)

	//s2 = append(s2, 8, 8, 4, 8)
	//s1[2] = 8
	a[4] = 100
	fmt.Println(s1) //[6  7  100]
	fmt.Println(s2) //	  	[5  9  8  8  4  8]
	println("a", a)
	println("s1", s1)// 这里s1是容量为5长度为3
	println("s2", s2)// 这里s2是容量为3长度为2
}
