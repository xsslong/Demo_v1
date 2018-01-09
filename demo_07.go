package main

import "fmt"

func main() {
	a := 8
	switch a {
	case 0:
		fmt.Println("a==0")
	case 1:
		fmt.Println("a==1")
	case 2:
		fmt.Println("a==2")
	case 3:
		fmt.Println("a==3")
	case 4:
		fmt.Println("a==4")
	case 5:
		fmt.Println("a==5")
	default:
		fmt.Println("a都不等于")
	}

	b := 3
	switch {
	case b > 0:
		fmt.Println("b>0")
		fallthrough
	case b > 1:
		fmt.Println("b>1")
	case b > 2:
		fmt.Println("b>2")
	case b > 3:
		fmt.Println("b>3")
	case b > 4:
		fmt.Println("b>4")
	case b > 5:
		fmt.Println("b>5")
	}

	c := 4
	switch {
	case c > 0:
		fmt.Println("c>0")
		fallthrough
	case c > 1:
		fmt.Println("c>1")
	case c > 2:
		fmt.Println("c>2")
	case c > 3:
		fmt.Println("c>3")
	case c > 4:
		fmt.Println("c>4")
	case c > 5:
		fmt.Println("c>5")
	}
	fmt.Println(c)
	fmt.Println(186.4 + 49.2 + 108.5 + 53.6 + 67.5 + 35.8 + 189.3 + 88 + 12.6)

LAB:
	for i := 1; i < 10; i++ {
		for {
			fmt.Println(i)
			continue LAB
		}
	}
	fmt.Println("执行完成")
}
