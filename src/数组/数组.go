package main

import "fmt"

/* 数组是具有相同唯一类型的一组已编号且长度固定的数据项序列 */
// 声明数组的语法格式
// var variable_name [Size] variable_type
var List1 [10]int
var List2 = [3]int{1, 2, 3}

// 如果数组长度不确定,可以用[...]代替数组的长度,编辑器会根据元素个数自行推断
var List3 = [...]int{0, 1, 2, 3, 4, 5, 6, 7}

// 如果设置了数组长度,可以按照下标来初始化,其索引从0开始
var List4 = [5]int{1: 2, 3: 4}

func main() {
	fmt.Println(List1, List2, List3, List4)
	fmt.Println(twoList)
	// 按照索引访问
	fmt.Println(twoList[0][1])
}

/* 多维数组 */
// var variable_name [size][size]...[size] variable_type
// 声明一个三维数组
var threeList [3][3][3]int

// 声明一个二维数组
var twoList = [2][2]int{{1, 2}, {2, 3}}
