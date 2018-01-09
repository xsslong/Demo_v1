package main

import (
	"strings"
	"fmt"
	//"path/filepath"
	//"os"
)

func main() {
	newstr := strings.Replace("xetadtqartatrt", "t", "O", -1)
	fmt.Println(newstr)


	//子串sep在字符串s中第一次出现的位置，不存在则返回-1。
	//func Index(s, sep string) int
	positionS := strings.Index("ASDWEFASDFAQWERF", "WE")
	fmt.Println(positionS)

	//子串sep在字符串s中最后一次出现的位置，不存在则返回-1。
	//func LastIndex(s, sep string) int
	positionL := strings.LastIndex("ASDWEFASDFAQWERF", "WE")
	fmt.Println(positionL)
}
