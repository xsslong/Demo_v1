package main

import "fmt"

type personD struct {
	id   int
	name string
}

func main() {
	StructTest06Base()
}

func StructTest06Base() {
	structTest0601()
	structTest0602()
}

//1.函数func, 接收值类型参数的函数
func valueIntTest(a int) int {
	return a + 10
}

//2.方法method, 接收指针类型参数的函数
func pointerIntTest(a *int) int {
	return *a + 10
}

func structTest0601() {
	a := 2
	fmt.Println("valueIntTest:", valueIntTest(a))

	b := 5
	fmt.Println("pointerIntTest:", pointerIntTest(&b))
}

// 接收者为值类型
func (p personD) valueShowName() {
	fmt.Println(p.name)
}

// 接收者为指针类型
func (p *personD) pointShowName() {
	fmt.Println(p.name)
}

func structTest0602() {
	//值类型调用方法
	personValue := personD{101, "Will Smith"}
	personValue.valueShowName()
	personValue.pointShowName()

	//指针类型调用方法
	personPointer := &personD{102, "Paul Tony"}
	personPointer.valueShowName()
	personPointer.pointShowName()
	//与普通函数不同，接收者为指针类型和值类型的方法，指针类型和值类型的变量均可相互调用
}
