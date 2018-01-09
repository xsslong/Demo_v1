package main

import "fmt"

var arr1 [2]int
var arr2 [2] int

var arr3 [3]string

var arr4 [3]bool

func main() {
	//arr2 = arr1
	//fmt.Print(arr1)
	//fmt.Print(arr2)
	//fmt.Print(arr3)
	//fmt.Print(arr4)
	//
	//arr5 := [2]float32{0.25, 5.54}
	//fmt.Print(arr5)
	//
	//var arr6 = [2]int{5, 6}
	//fmt.Print(arr6)
	//
	//arr7 := [5]int{3, -1}
	//fmt.Print(arr7)
	//
	//arr8 := [5]int{2: 666, 4: 86}
	//fmt.Print(arr8)
	//
	//arr9 := [...]string{"afasf", "shenm", 6: "我也不知道"}
	//fmt.Print(arr9)

	arr10 := [...]int{19: 88}
	//var p *[20]int = &arr10
	//var p = &arr10
	p := &arr10
	p[3] = 4
	fmt.Println(p)

	x, y := 35, 67
	m := [2]*int{&x, &y}
	fmt.Println(m)

	a1, a2 := [2]int{5}, [2]int{3, 5}
	fmt.Println(a1 == a2)

	q1 := [2][3][5][4]int{}
	fmt.Println(q1)

	qwe := [...]int{41, 34, 56, 23, 67, 38}
	fmt.Println(qwe, "排序前")

	l := len(qwe)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if qwe[i] > qwe[j] {
				temp := qwe[j]
				qwe[j] = qwe[i]
				qwe[i] = temp
			}
		}
	}
	fmt.Println(qwe, "排序后")
}
