package main

import "fmt"

func main() {
	//	s := "BrainWu"
	//	if v, ok := interface{}(s).(string); ok {
	//		fmt.Println(v)
	//	}
	panduan(Handler(ServeHTTP))
}

func ServeHTTP(s string) {
	fmt.Println(s)
}

type Handler func(string)

func panduan(in interface{}) {
	//Handler(in)("wujunbin")
	if v, ok := in.(Handler); ok {
		//跟什么类型判断就只能调用什么类型的方法
		v("BrainWu")
	}
}
