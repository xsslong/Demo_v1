package main

import (
	"fmt"
	//"math"
)

func main() {
	fmt.Println(max(7876, 4566, 6124), "最大值")

	f := func(ww int) int {
		ww += 4
		return ww
	}

	fmt.Println(f)
	fmt.Println(f(16))

	//ff := mix(3, 5)
	//fmt.Println(ff(4))

	fmt.Println(mix(3, 5)(5))
}

// 函数返回三个数的最大值
func max(num1, num2, num3 int) int {
	if num1 > num2 {
		if num1 > num3 {
			return num1
		}
		return num3
	} else {
		if num2 > num3 {
			return num2
		}
		return num3
	}
	//math.Max()
}

// 先求和再求积
func mix(num1, num2 int) func(i int) (result int) {
	return func(i int) (result int) {
		return (num1 + num2) * i
	}
}
