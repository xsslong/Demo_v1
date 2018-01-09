package main

import (
	"fmt"
	//"os"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, 2145, "qwert")
	fmt.Fprintln(os.Stdout, 9527, "yyy", false)
	fmt.Fprintf(os.Stdout, "8848", true)

	//fmt.Sprint("Sprint", 'q', 110, false)
	//fmt.Sprintln("Sprintln", 'w', 1220, true)
	//fmt.Sprintf("Sprintf", 'e', 140, false)

	//fmt.Print("Print", 123, true, 'a')
	//fmt.Println("Println", 25445, true, 'b')
	//fmt.Printf("Printf", 1214, false, "金戈铁马")
}
