package main

import "fmt"

func main() {
	b := [3]int{1, 2, 3}
	f2(&b)
	fmt.Println(b)// [11, 22, 33]
}

func f2(b *[3]int) {
	b[0] = 11
	b[1] = 22
	b[2] = 33
}
