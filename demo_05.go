package main

import "fmt"

var a, b = 6, 12 // 4 = 0110, 6 = 1010

const (
	B  = 1 << (iota * 10)
	KB
	MB
	GB
	TB
)

func main() {
	//fmt.Println(^a)
	fmt.Println(a ^ b)
	fmt.Println(1 << 10)
	fmt.Println(1 << 10 << 10)
	fmt.Println(1 << 10 << 10 >> 10)

	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)

	//var f *int = &a
	var f = &a
	fmt.Println(f)
	fmt.Println(*f)

	// go语言中++、--智能作为单独语句运算, 不能作为变量
	//fmt.Println(a++)// 编译错误
	a++
	fmt.Println(a)

}
