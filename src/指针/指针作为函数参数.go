package main

import "fmt"

func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	var a int = 1
	var b int = 2

	fmt.Println("a,b:", a, b)

	// 交换 a,b 的值
	swap(&a, &b)
	fmt.Println("a,b:", a, b)
}
