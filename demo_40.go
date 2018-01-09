package main

import "fmt"

//func main() {
//	slice1 := []int{1,2,3,4,5}
//	// TODO: 注意,下面这种slice是共享内存的,并不是copy
//	slice2 := slice1[2:4]
//	fmt.Println(slice1, slice2)// [1,2,3,4,5], [3, 4]
//	slice2[0] = 333
//	// 输出1 2 333 4 5
//	fmt.Println(slice1)// [1,2,333,4,5]
//}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[2:4] // [3, 4]
	slice2 = append(slice2, 66)// [3, 4, 66], slice1中的5会被66覆盖
	fmt.Println(slice1, slice2) // [1, 2, 3, 4, 66], [3, 4, 66]
}
