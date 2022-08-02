package main

import "fmt"

/* 切片 */
// 与数组相比,其长度不固定

func main() {
	// 定义切片
	// var identifier []type	切片不需要说明长度
	// 或者使用 make() 来创建切片
	// var slice1 []type = make([]type, len)	len 指代切片的初始长度
	// slice1 := make([]type, len)

	// 切片的初始化
	s := []int{1, 2, 3}
	var s2 = make([]int, 3, 5) // (type,len,cap)

	printSlice(s)  // 3 3 [1 2 3]
	printSlice(s2) // 3 5 [0 0 0]

	/* 空切片 */
	// 切片在未初始化之前默认为 nil, 长度为 0
	var s3 []int
	printSlice(s3) // 0 0 []

	/* 切片的截取 */
	// 从数组中获取一个切片
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8}
	// s := arr[start:end]	左闭右开
	s4 := arr[0:5]
	printSlice(s4) // 5 10 [1 2 3 4 5]
	s5 := s4[1:3]
	printSlice(s5) // 2 9 [2 3]

	/* append() 和 copy() */
	s5 = append(s5, -1)
	printSlice(s5) // 3 9 [2 3 -1]
	s5 = append(s5, 1, 2, 3)
	printSlice(s5) // 6 9 [2 3 -1 1 2 3]

	// copy() 详细再写一个
	var s6 []int = make([]int, 5) // 5 5 [0 0 0 0 0]
	copy(s6, s5)
	printSlice(s6) // 5 5 [2 3 -1 1 2]

}

// len() 测量长度 (实际)
// cap() 测量切片最长可以达到多少 (容量)
// len 切片 <=cap 切片 <=len 数组
func printSlice(s []int) {
	fmt.Println(len(s), cap(s), s)
}

// 切片删除
// 没有直接删除的方法
func test1() {
	s := []int{1, 2, 3, 4, 5}
	// 删除 3
	s = append(s[:2], s[3:]...)
	fmt.Println(s) // [1, 2, 4, 5]
}

// 切片更新
func test2() {
	s := []int{1, 2, 3, 4}
	s[1] = 100
	fmt.Println(s) // [100 2 3 4]
}
