package main

import "fmt"

// init
// 先于 main 函数自动执行,不可被调用 实现包级别的初始化操作
// 可以有多个 init 函数

// 顺序: 变量初始化 -> init -> main

var i int = initVar()

func init() {
	fmt.Println("init...")
}

func init() {
	fmt.Println("init2...")
}

func initVar() int {
	fmt.Println("100")
	return 100
}
func main() {
	fmt.Println("main...")
}
