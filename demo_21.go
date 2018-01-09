package main

import "fmt"

type AStudent struct {
	Name string
}

type BStudent struct {
	Name string
}

func main() {
	a := AStudent{
		"Lily",
	}

	b := BStudent{
		"Bobo",
	}
	a.printName()
	b.printName()
	fmt.Println(a)
	fmt.Println(b)
}

func (a *AStudent) printName() {
	a.Name = "LiliLily"
	fmt.Println(a.Name)
}

func (b BStudent) printName() {
	b.Name = "BoboBobo"
	fmt.Println(b.Name)
}
