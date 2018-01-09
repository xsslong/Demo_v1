package main

import (
	"flag"
	"fmt"
)

func main() {
	//说明:
	//	像flag.Int、flag.Bool、flag.String这样的函数格式都是一样的，调用的时候需要传入3个参数
	//参数的说明如下:
	//	第一个arg表示参数名称，在控制台的时候,提供给用户使用.
	//	第二个arg表示默认值，如果用户在控制台没有给该参数赋值的话,就会使用该默认值.
	//	第三个arg表示使用说明和描述,在控制台中输入-arg的时候会显示该说明, 类似-help

	married := flag.Bool("married", false, "Are you married?")
	age := flag.Int("age", 22, "How old are you?")
	name := flag.String("name", "", "What your name?")

	var address string
	//flag.StringVar这样的函数第一个参数换成了变量地址，后面的参数和flag.String是一样的。
	flag.StringVar(&address, "address", "GuangZhou", "Where is your address?")

	flag.Parse() //解析输入的参数

	fmt.Println("输出的参数married的值是:", *married) //不加*号的话,输出的是内存地址
	fmt.Println("输出的参数age的值是:", *age)
	fmt.Println("输出的参数name的值是:", *name)
	fmt.Println("输出的参数address的值是:", address)
}

//https://www.cnblogs.com/sunailong/p/7866660.html
