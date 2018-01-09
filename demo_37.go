package main

import (
	"fmt"
	"time"
)

func main() {
	fde()
	fmt.Println("main end")
}

func fde() {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("defer func start")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，"bug"
		}
		fmt.Println("defer func end")
	}()
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")

	fmt.Println("1")
	a := []string{"a", "b"}
	fmt.Println(a[3]) // 越界访问，肯定出现异常
	panic("buggggg")  // 上面已经出现异常了, 所以从此以下代码不会再执行
	fmt.Println("4")  // 不会运行的.
	time.Sleep(1 * time.Second)
}
