package main

import (
	"fmt"
	"strconv"
)

// 常量的定义
const q = 5
const w int = 88

// 采用组的形式定义
//const (
//	t bool = false
//	y = 5.2
//	r = "退款"
//	u = len(r)
//)

// 同时定义多个常量
const t, y, r, u = false, 5.2, "搞笑", len(r)

//const (
//	t, y, r, u = false, 5.2, "搞笑", len(r)
//)

var sss = "oppo"

const (
	o = 55
	k
	j
	//l = len(sss) // 编译不通过, =右侧必须是常量的表达式
	//l = strconv.Itoa(o) // 编译不通过, =右侧必须是常量的表达式

	m, n = 87, 52.36
	v, i
)

const (
	aaaa  = "aa"  // iota = 0
	aaaaa = "eaa" // iota = 1
	bbbb  = iota  // iota = 2
	cccc          // iota = 3
	dddd          // iota = 4
	eeee          // iota = 5
)

const (
	ffff = iota // iota = 0
	gggg        // iota = 1
)

func main() {
	var a = 68
	fmt.Println(a)

	b := string(a)
	fmt.Println(b)

	c := strconv.Itoa(a)
	fmt.Println(c)

	d, error := strconv.Atoi("8869")
	fmt.Println(d)
	fmt.Println(error)

	fmt.Println(t)
	fmt.Println(y)
	fmt.Println(r)
	fmt.Println(u)

	fmt.Println(o)
	fmt.Println(k)
	fmt.Println(j)

	fmt.Println(v)
	fmt.Println(i)

	fmt.Println(bbbb)
	fmt.Println(cccc)
	fmt.Println(dddd)
	fmt.Println(eeee)

	fmt.Println(ffff)
	fmt.Println(gggg)
}
