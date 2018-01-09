package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 1,定义切片
	//s1, s2, s3 := []int, []string, []bool
	//var s1 [2]int
	//var s2 [4]string
	//var s3 []bool
	//fmt.Println(s1)
	//fmt.Println(s2)
	//fmt.Println(s3)

	// 2,通过make函数来创造切片
	//tt := make([]string, 3, 9)
	//fmt.Println(tt)
	//fmt.Println(len(tt))
	//
	//// 3,直接通过类似数组的赋值方式来初始化切片
	//pp := make([]string, 4, 5)
	//jj := []int{2, 5, 6, 8, 9}// 其中其cap=len=5
	//
	//pp[2] = "钛金手机"
	//fmt.Println(pp)
	//fmt.Println(len(pp))
	//fmt.Println(jj)
	//fmt.Println(len(jj))

	// 通过数组来得到切片
	//oo := [...]int{2, 3, 5, 7, 12, 4, 6, 9, 0, 3}
	//ss1 := oo[2:5]
	//ss2 := oo[3:10]
	//ss3 := oo[4:]
	//ss4 := oo[:6]
	//ss5 := oo[5]
	//ss6 := oo[:]
	//
	//fmt.Println(oo)  // [2 3 5 7 12 4 6 9 0 3]
	//fmt.Println(ss1) // [5 7 12]
	//fmt.Println(ss2) // [7 12 4 6 9 0 3]// [startIndex:endIndex] endIndex最大取值是len
	//fmt.Println(ss3) // [12 4 6 9 0 3]
	//fmt.Println(ss4) // [2 3 5 7 12 4]
	//fmt.Println(ss5) // 4 ss5就是个int型数据
	//fmt.Println(ss6) // [2 3 5 7 12 4 6 9 0 3]
	//
	//fmt.Println("%p\n", oo)  //
	//fmt.Println("%p\n", ss6) //

	// 通过切片来得到切片
	s := []int{2, 3, 5, 7, 12, 4, 6, 9, 0, 3}
	sss1 := s[2:5]
	sss2 := s[3:10]
	sss3 := s[4:]
	sss4 := s[:6]
	sss5 := s[5]
	sss6 := s[:]

	fmt.Println(s)    // [2 3 5 7 12 4 6 9 0 3]
	fmt.Println(sss1) // [5 7 12]
	fmt.Println(sss2) // [7 12 4 6 9 0 3]// [startIndex:endIndex] endIndex最大取值是len
	fmt.Println(sss3) // [12 4 6 9 0 3]
	fmt.Println(sss4) // [2 3 5 7 12 4]
	fmt.Println(sss5) // 4 ss5就是个int型数据
	fmt.Println(sss6) // [2 3 5 7 12 4 6 9 0 3]
	fmt.Println("s:", reflect.TypeOf(s))
	fmt.Println("sss5:", reflect.TypeOf(sss5))
	fmt.Println("sss6:", reflect.TypeOf(sss6))

	println("sss1", sss1)
	println("sss2", sss2)
	println("sss3", sss3)
	println("sss4", sss4)
	println("sss5", sss5)
	println("sss6", sss6)

	//s := []int{2, 3, 5, 7, 12, 4, 6, 9, 0, 3}
	//fmt.Println(s)
	//fmt.Println(len(s))
	//fmt.Println(cap(s))
	//
	//s1 := make([]string, 4, 6)
	//s1[3] = "aa"
	//fmt.Println(s1)
	//fmt.Println(len(s1))
	//fmt.Println(cap(s1))
	//
	//var numbers = make([]int, 3, 5)
	//printSlice(numbers)
	//
	//var numbers1 []int
	//printSlice(numbers1)
	//if numbers1 == nil {
	//	fmt.Printf("切片是空的")
	//}

	//st := make([]int, 4, 6)
	////st := []int{4,5}
	//fmt.Println(st)
	//println(st)
	//fmt.Println("st:", reflect.TypeOf(st))
	//
	//stt := append(st, 34)
	//fmt.Println(stt)
	//println(stt)
	//fmt.Println("stt:", reflect.TypeOf(stt))
	//
	//sttt := append(stt, 46, 27)
	//fmt.Println(sttt)
	//println(sttt)
	//fmt.Println("sttt:", reflect.TypeOf(sttt))
	//
	//
	//fmt.Println("Hello world!")
	//println("Hello world!")
}

//func printSlice(x []int) {
//	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
//}
