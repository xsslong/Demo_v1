package main

import "fmt"

type Animal interface {
	Grow()
	Move(string) string
}

type Cat struct {
	Name    string
	Age     int
	Address string
}

func (cat *Cat) Grow() {

}

func (cat *Cat) Move(s string) string {
	return cat.Name
}

func main() {
	myCat := Cat{"Little C", 2, "In the house"}
	animal, ok := interface{}(&myCat).(Animal)
	fmt.Printf("%v, %v\n", ok, animal)
}
