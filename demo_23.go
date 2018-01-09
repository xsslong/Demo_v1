package main

import (
	"fmt"
	"reflect"
)

type Phone interface {
	call()
}

type IPhone struct {
}

type XiaoMi struct {
}

func main() {
	//iPhone := IPhone{}
	//iPhone.call()
	//
	//xiaoMi := XiaoMi{}
	//xiaoMi.call()

	var iPhone Phone
	iPhone = new(IPhone)
	iPhone.call()
	fmt.Println(reflect.TypeOf(iPhone))
}

func (iPhone IPhone) call() {
	fmt.Println("I am IPhone")
}

func (xiaoMi XiaoMi) call() {
	fmt.Println("I am XiaoMi")
}
