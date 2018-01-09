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

type people2 struct {
	Name   string
	Age    int
	Class  string
	WorkId string
}

type teacher2 struct {
	Pp      people2
	Subject string
	WorkId  string
}

type student2 struct {
	people2
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

	t := teacher2{
		Pp: people2{
			Name: "张小泉",
			Age:  36,
		},
		Subject: "英语",
		WorkId:  "9527",
	}

	s := student2{
		people2: people2{
			Name: "王二嘎子",
			Age:  16,
		},
		Class: "三年二班",
		Grade: 3,
	}
	t.WorkId = "3306"
	t.Pp.WorkId = "10010"
	t.Pp.Name = "张大泉"
	t.Subject = "体育"

	//s.people2.Name = "王嘎子"
	s.Name = "王嘎子"
	s.Class = "五班"
	s.people2.Class = "七班"
	s.Grade = 4
	fmt.Println(t)
	fmt.Println(s)
}
