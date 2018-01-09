package main

import "fmt"

type Rect struct {
	x, y float32
}

func main() {
	rect := Rect{
		5.54,
		6.35,
	}
	fmt.Println("方法:", rect.area())
	fmt.Println("函数:", area(rect))
}

func (r Rect) area() (s float32) {
	s = r.x * r.y
	return
}

func area(r Rect) (s float32) {
	s = r.x * r.y
	return
}
