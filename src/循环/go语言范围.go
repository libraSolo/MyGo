package main

import "fmt"

/* Range */
// for 迭代 array, slice, channel, map 中
// array 和 slice 返回索引-值,在 map 中返回 key-value

func main() {
	// 简单循环
	nums := []int{1, 2, 3, 4}
	length := 0
	for range nums {
		length++
	}
	fmt.Println(length)

}
