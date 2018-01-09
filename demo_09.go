package main

import "fmt"

func main() {
	// 输出1~100以内的素数
LAB:
	for i := 3; i <= 100; i++ {
		for j := 2; j <= i-1; j++ {
			if i%j == 0 {
				continue LAB
			}
		}
		fmt.Print(i, " ")
	}
	//fmt.Print(5.0/4)

	arr := []int{8, 4, 16, 27, 6, 33, 6, 74, 50}
	fmt.Println(getAverage(arr), "平均值")

	xxxx, yyyy, zzzz := sort(4, 7, 8)
	fmt.Println(xxxx, "求和1")
	fmt.Println(yyyy, "求和2")
	fmt.Println(zzzz, "求和3")
}

// 求一个数组的平均值
func getAverage(arr []int) float32 {
	size := len(arr)
	sum := 0
	for i := 0; i < size; i++ {
		sum += arr[i]
	}
	sum1 := float32(sum)
	size1 := float32(size)
	return sum1 / size1
}

// 求和
func sort(x, y, z int) (a, b, c int) {
	a = x + y
	b = x + z
	c = y + z
	return
}