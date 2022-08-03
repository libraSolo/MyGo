package main

import "fmt"

func main() {
	// 声明值
	var a int
	// 声明指针
	var ptr *int
	// 声明指针的指针
	var pptr **int

	a = 1
	ptr = &a
	pptr = &ptr
	fmt.Println(a, ptr, pptr)
}
