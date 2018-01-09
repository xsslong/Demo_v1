package main

import "fmt"

func main() {
	if i := 1; i < 10 {
		fmt.Println("小于10")
	}

	if i, p := 12, 4; i*p < 50 {
		fmt.Println("小于10")
	}

	//方式1: 无条件语句, 无限循环(最后一定要跳出循环)
	a := 6
	for {
		a++
		if a <= 10 {
			fmt.Println(a, "小于10的a")
		} else {
			break
		}
	}

	//方式2: 一个条件语句
	b := 4
	for b <= 11 {
		b++
		fmt.Println(b, "5到12")
	}

	//方式2: 经典的三个条件语句
	for c := 2; c <= 20; c++ {
		fmt.Println(c, "2到20")
	}
}
