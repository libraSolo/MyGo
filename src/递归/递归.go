package main

import "fmt"

/* 递归 */

// 递归求阶乘
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

// 斐波那契数列
func fibonacci(n int) (result int) {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	var i uint64 = 3
	fmt.Println(Factorial(i))
	fmt.Println(fibonacci(7))
}
