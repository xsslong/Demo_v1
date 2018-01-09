package main

import "fmt"

func main() {

	// 根据new的定义
	//a := new(int)
	a := newInt()
	fmt.Println(a)
	fmt.Println(*a)
}

func newInt() *int {
	var a int
	return &a
}
