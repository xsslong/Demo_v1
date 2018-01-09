package main

import "fmt"

func main() {
	dict := make(map[string]int)
	dict["AAA"] = 333
	dict["BBB"] = 666
	dict["CCC"] = 999
	// map返回值有两种情况,第一种是仅仅返回值,第二是返回值和exist-bool值
	d1 := dict["AAA"]
	fmt.Println(d1)

	// 通过第二个参数判断这个key是否存在
	d2, exist := dict["AAA"]
	fmt.Println(exist, d2)
}
