package main

import "fmt"

func main() {
	var array3 [3]*string

	array4 := [3]*string{new(string), new(string), new(string)}
	*array4[0] = "Red"
	*array4[1] = "Blue"
	*array4[2] = "Green"
	// 赋值OK
	array3 = array4

	// 大小错误
	var array31 [4]*string
	array31 = array4

	// 类型错误
	var array32 [3]string
	array32 = array4

	fmt.Println(array3)
	fmt.Println(array31)
	fmt.Println(array32)
	fmt.Println(array4)
}
