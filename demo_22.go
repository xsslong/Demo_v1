package main

import "fmt"

type INT int

func main() {
	var a INT // 这里不能使用 a:=0, 需要明确其为INT类型
	a.Increase()
	//(&a).Increase()
	fmt.Println("a最终值: ", a)
}

func (i *INT) Increase() {
	for j := 0; j < 100; j++ {
		fmt.Println(*i)
		*i += 1
	}
	fmt.Println(*i)
}
