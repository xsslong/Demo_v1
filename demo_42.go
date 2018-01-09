package main

import (
	"fmt"
	"strings"
	"time"
	"net/http"
)

func main() {
	s := "/aaapao"

	// HasPrefix(s, prefix string) bool 判断字符串是否有前缀
	fmt.Println(strings.HasPrefix(s, "/"))
	fmt.Println(strings.HasPrefix(s, "\\"))

	// HasSuffix(s, suffix string) bool 判断字符串是否有后缀
	fmt.Println(strings.HasSuffix(s, "ao"))
	fmt.Println(strings.HasSuffix(s, "/"))
	fmt.Println(time.Now().Add(3600000000000).Unix())


	resp, _ := http.Get("http://www.baidu.com")
	fmt.Println(resp)
}
