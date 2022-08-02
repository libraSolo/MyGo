package main

import "fmt"

// 语法格式
//func ([参数])[返回值]

func main() {
	// 匿名函数
	sum := func(a, b int) int {
		return a + b
	}
	res := sum(1, 2)
	fmt.Println(res)

	// 直接调用
	r := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(r)
}
