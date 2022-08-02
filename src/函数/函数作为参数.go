package main

import "fmt"

// 声明一个函数类型
type cb func(int) int

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}

func test1(x int) int {
	return x + 1
}

func main() {
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})

	// 声明一个函数类型
	type f1 func(int) int
	var ff f1
	ff = test1
	res := ff(1)
	fmt.Printf("%v", res)

}
