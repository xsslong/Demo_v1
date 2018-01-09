package main

import "fmt"

// 类型的别名
type (
	zx int
	fd32 float32
	文本 string
)

var (
	// 直接定义变量
	a int
	b float32
	c string

	// 通过类型的别名定义变量
	d zx
	e fd32
	f 文本
)

var (
	aa bool

	bb byte
	cc int
	dd float32
	ee complex64
	ff uintptr

	gg 文本

	kk []int
	ll [3]int
)

func main() {
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)

	fmt.Println(ee)
	fmt.Println(ff)

	fmt.Println(gg)

	fmt.Println(kk)
	fmt.Println(ll)

	// go语言中数据类型不存在隐式转换, 所有的类型转换必须显示声明
	var v float32 = 65.234
	v_ := int(v)
	fmt.Println(v_)

	// 转换只能发生在两者相互兼容类型之间
	//var c bool = true
	//c_ := int(c)

	var qq = 50
	qq_ := string(qq)
	fmt.Println(qq_)
}
