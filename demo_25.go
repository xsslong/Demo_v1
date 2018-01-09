package main

import "fmt"

//type Integer int

//type comperAdd interface {
//	comper(b Integer) bool
//	add(b Integer)
//}
//
//func main() {
//	var a Integer = 8848
//	var b1 comperAdd = &a
//	//var b2 comperAdd = a // 类型不兼容, Integer实现了comperAdd接口, 根据接口的定义, 接口是其实现了的
//	fmt.Println(b1)
//}
//
//func (a Integer) comper(b Integer) bool {
//	return a > b
//}
//
//func (a *Integer) add(b Integer) {
//	*a += b
//}

type USB interface {
	USBType() string
	Connect()
}

type MyPhone struct {
	Name string
}

type MyUpan struct {
	Name string
	Cpa  int
}

func main() {
	//var myPhone USB
	myPhone := MyPhone{
		Name: "XiaoMi",
	}
	myPhone.USBType()
	myPhone.Connect()
	DisConnect(myPhone)

	//var myUpan USB
	myUpan := MyUpan{
		Name: "Kingston",
	}
	//myUpan.USBType()
	myUpan.Connect()
	//DisConnect(myUpan)
}

func (myPhone MyPhone) USBType() string {
	fmt.Println("My USBType is", myPhone.Name)
	return myPhone.Name
}

func (myPhone MyPhone) Connect() {
	fmt.Println("Phone Connectted")
}

//func (myUpan MyUpan) USBType() string {
//	fmt.Println("My USBType is", myUpan.Name)
//	return myUpan.Name
//}

func (myUpan MyUpan) Connect() {
	fmt.Println("Upna Connectted")
}

func DisConnect(usb USB) {
	//if usbType, ok := usb.(MyPhone); ok {
	//	fmt.Println(usbType.Name, "*******DisConnect Success******")
	//}

	switch usbType := usb.(type) {
	case MyPhone:
		fmt.Println(usbType.Name, "*******DisConnect Success******")
	default:
		fmt.Println("*******UNKONW DEVICE******")
	}
}
