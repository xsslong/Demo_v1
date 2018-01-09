package main

import "fmt"

//import "strings"// 常规调用
//import StrUtil "strings"// 别名调用
import . "strings" // 别名为. 省略调用

type alias string

//var (
//	q = 90
//	w = false
//	r = 0.77
//	t = "这他妈是什么"
//)

func main() {

	//var b string 这里声明b是string数据类型
	var a alias // 这里的alias是一个string类型的别名, 即alias就是代表string类型的书写(不建议使用别名)
	a = "这真的是一个文本"
	fmt.Print(a)

	//var b int // 先声明后赋值
	//b = 1.2   // 当b被赋值为非int类型数据, 可通过编译, 但输出报错
	//fmt.Print(b)

	var c float32 = 8.6 // 直接声明赋值, 赋值编译, 自动检测数据类型, 可将float32省略 var c = 8
	fmt.Print(c)

	//var d = 777
	d := 777 // 简写赋值的方式, 去掉var自动编译检测数据类型
	var e = false
	f := true
	g := "大王叫我来巡山"
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// 全局变量声明后未使用编译通过, 只是会提示未使用; 局部变量声明之后未使用则无法通过编译
	//var (
	//	q = 90
	//	w = false
	//	r = 0.77
	//	t = "这他妈是什么"
	//)
	//fmt.Println(q)
	//fmt.Println(w)
	//fmt.Println(r)
	//fmt.Println(t)

	//var x, y, z, u int = 8, 8, 4, 8
	//var x, y, z, u = 8, 8, 4, 8
	x, y, z, u := 8, 8, 4, 8
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(u)

	// 局部变量才能使用:=的方式声明, 全局变量不可以使用:=的方式声明
	p, o, i := true, 8, "钛金手机"
	fmt.Println(p)
	fmt.Println(o)
	fmt.Println(i)

	// 常规调用
	//var compare = strings.Compare("A","F")
	//fmt.Println(compare)

	// 别名调用
	//var compare = StrUtil.Compare("K", "H")
	//fmt.Println(compare)

	// 省略调用
	var compare = Compare("K", "H")
	fmt.Println(compare)

	//fmt.Println(HR*5)

}
