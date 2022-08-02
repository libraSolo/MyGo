package main

import "fmt"

// 声明全局变量
var g int

func main() {
	// 声明局部变量
	var a, b, c int

	// 初始化参数
	a = 10
	b = 20
	c = a + b

	fmt.Println(a, b, c)
}
