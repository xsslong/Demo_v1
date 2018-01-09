package main

import "fmt"

func main() {

	//s1 := []int{6, 72, 34, 57, 89}
	//s2 := []int{11, 45, 23}
	//ss1 := copy(s1, s2) // 返回被copy元素的个数
	//fmt.Println("被copy元素个数", ss1)
	//fmt.Println("s1", s1)
	//fmt.Println("s2",s2)
	//println("s1", s1)
	//println("s2", s2)

	//s1 := []int{6, 72, 34, 57, 89}
	//s2 := []int{11, 45, 23, 3, 45, 672, 1356}
	//ss1 := copy(s1, s2) // 返回被copy元素的个数
	//fmt.Println("被copy元素个数", ss1)
	//fmt.Println("s1", s1)
	//fmt.Println("s2",s2)
	//println("s1", s1)
	//println("s2", s2)

	// 部分拷贝
	s1 := []int{6, 72, 34, 57, 89, 33, 45, 66}
	s2 := []int{11, 45, 23, 22, 59, 64}
	//ss1 := copy(s1[2:], s2) // 返回被copy元素的个数
	//fmt.Println("被copy元素个数", ss1)
	//fmt.Println("s1", s1)
	//fmt.Println("s2", s2)
	//println("s1", s1)
	//println("s2", s2)
	for _, value := range s1 {
		//fmt.Print("key:", key)
		fmt.Println(" value:", value)
	}
	//	fmt.Print("key:", key)
	//	fmt.Println(" s1:", s1[key])
	//}
}
