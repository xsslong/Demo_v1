package main

import (
	"path/filepath"
	"os"
	"fmt"
	"strings"
	"log"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(dir)
	if err != nil {
		log.Fatal("Get Root Path Error:", err)
	}
	dirctory := strings.Replace(dir, "\\", "/", -1)

	fmt.Println(dirctory)

	runes := []rune(dirctory)
	fmt.Println(runes)
	l := 0 + strings.LastIndex(dirctory, "/")
	fmt.Println(l)
	if l > len(runes) {
		l = len(runes)
	}
	fmt.Println(string(runes[0:l]))
	fmt.Println(os.Hostname())
}
