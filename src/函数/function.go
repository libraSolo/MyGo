package main

import (
	"fmt"
	"math"
)

// 返回两个数中,较大的数
func max(num1, num2 int) int {

	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

// 按照从小到大返回(返回多个数)
func sort(num1, num2 int) (int, int) {
	if num1 < num2 {
		return num1, num2
	} else {
		return num2, num1
	}
}

// 一般情况下,调用函数使用的是传递值,不影响实际参数
func swap(num1, num2 int) int {
	var temp int
	temp = num1
	num1 = num2
	num2 = temp
	return temp
}

// 引用传递, 将影响实际参数
func swap2(num1, num2 *int) {
	var temp int
	temp = *num1
	*num1 = *num2
	*num2 = temp
}

// 传递任意数量的值
func test1(args ...int) {
	for _, i := range args {
		fmt.Println(i)
	}
}

func main() {
	var a = 1
	var b = 2
	fmt.Println(max(1, 2))
	fmt.Println(sort(2, 1))

	swap(a, b)        // 传递的值
	fmt.Println(a, b) // 1, 2 此时的 a, b 并没有交换变量

	swap2(&a, &b)     // 传递 a,b 的地址
	fmt.Println(a, b) // 2, 1 此时的 a, b 交换了变量

	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(9))
}
