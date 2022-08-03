package main

import "fmt"

func main() {
	// 类型定义
	type myInt int
	var i myInt
	i = 100
	fmt.Printf("%T\n", i)

	type myInt2 = int
	var i2 myInt
	i2 = 100
	fmt.Printf("%T\n", i2)

}
