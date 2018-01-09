package main

import "fmt"

type ppxiao struct {
	Name string
	Age  int
	Contact struct {
		ID, City string
		Phone    int
	}
}

func main() {
	huluwa := &ppxiao{
		Name: "金刚",
		Age:  188,
		Contact: struct {
			ID, City string
			Phone    int
		}{ID: "9527", City: "东莞", Phone: 110,},
	}

	huluwa.Name = "闪电"
	huluwa.Contact.City = "东京热"
	modifyPhone(huluwa)
	fmt.Println("葫芦娃:", huluwa)
	fmt.Println("葫芦娃Contact:", huluwa.Contact)
}

func modifyPhone(pp *ppxiao) {
	pp.Contact.Phone = 114
}
