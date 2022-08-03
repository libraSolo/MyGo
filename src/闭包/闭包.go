package main

// 闭包 = 函数 + 引用环境

import "fmt"

func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	f := add()
	r := f(1)
	fmt.Println(r)
	r = f(2)
	fmt.Println(r)
	r = f(3)
	fmt.Println(r)
	fmt.Println("---------")
	f1 := add()
	fmt.Println(f1(1))
	fmt.Println(f1(2))
}
