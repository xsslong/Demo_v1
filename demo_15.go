package main

import (
	"fmt"
	//"math"
	"reflect"
)

// 这个person是struct的名字
type person struct {
	Name string
	age  int
}

func main() {

	//m := make(map[int]map[string]string)
	//m[1] = make(map[string]string)
	//m[1]["你算"] = "哪根葱"
	//fmt.Println(m)
	//fmt.Println(m[1])
	//fmt.Println(m[1]["你算"])
	//fridend := &person{
	//	Name: "Lily",
	//	age:  19,
	//}
	////fmt.Println(fridend)
	//fmt.Println("原来的: ", fridend)
	//fridend.Name = "paipai"
	//modifyAge(fridend)
	//fmt.Println("修改后: ", fridend)
	//fmt.Println("TypeOf:", reflect.TypeOf(fridend))

	huluwa := &struct {
		Name string
		Age  int
	}{
		Name: "大娃",
		Age:  8,
	}
	huluwa.Age = 18
	fmt.Println("葫芦娃", huluwa)
}

func modifyAge(p *person) {
	p.age += 5
	fmt.Println("传入的: ", p)
}
