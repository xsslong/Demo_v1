package main // 只有package名称位main的包可以包含main函数, 一个可执行程序有且仅有一个main包

import "fmt"
import "time"

//import ss "fmt"// ss称为package别名, 可采用ss.调用
//import . "fmt"// . 做别名可以直接调用, 不建议使用, 容易误解
// .调用和别名不可同时使用

// 多个包的导入通过()的形式完成
//import (
//	"bufio"
//	"strings"
//	"syscall"
//	"time"
//
//	"github.com/elazarl/goproxy"
//
//	"github.com/yinghuocho/i18n"
//	"github.com/yinghuocho/tarfs"
//)

// 通过const 来定义常量
//const PI = 6
//const HR = 8585

const (
	PI = 6
	HR = 8585
)

// 通过var 来定义定义和申明全局变量
//var name = "圆周长"
//var length = 4
//var aa = 5

var (
	name   = "圆周长"
	length = 4
	aaa     = 5
)

// 一般类型的声明,  type + 名称 + 类型、结构、接口
type newType int

// 结构的声明
type gopher struct {
}

// 接口的声明
type golang interface {
}

// 简写模式
type (
	newType_ int
	gopher_ struct{}
	golang_ interface{}
)

func main() {
	// 导入包后, 可以使用<PackageName>.<FuncName>来对包中的函数进行调用
	// go语言是强编译语言, 如果导入包后未调用其中的函数或者类型将会报出编译错误
	fmt.Printf(name + string(PI*length))
	fmt.Printf("我没有钱" + time.ANSIC)
	//fmt.Printf("我没有钱" + fmt.commaSpaceString)
	//fmt.Printf("我没有钱" + time.stdMonth)
}
