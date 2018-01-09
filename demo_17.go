package main

import "fmt"

//type car struct {
//	Name  string
//	Price float32
//}
//
//type car2 struct {
//	Name  string
//	Price float32
//}

type people struct {
	Name string
	Age  int
}

type teacher struct {
	Pp      people
	Subject string
	WorkId  string
}

type student struct {
	people
	Class string
	Grade int
}

func main() {
	//c := car{
	//	"BMW",
	//	40.4,
	//}
	//
	//d := car2{
	//	"BMW",
	//	40.4,
	//}
	//fmt.Println(c == d) // 不同的结构体不能进行比较

	t := teacher{
		Pp: people{
			Name: "张小泉",
			Age:  36,
		},
		Subject: "英语",
		WorkId:  "9527",
	}

	s := student{
		people: people{
			Name: "王二嘎子",
			Age:  16,
		},
		Class: "三年二班",
		Grade: 3,
	}
	t.Pp.Name = "张大泉"
	t.Subject = "体育"

	//s.people.Name = "王嘎子"
	s.Name = "王嘎子"
	s.Grade = 4
	fmt.Println(t)
	fmt.Println(s)
}
