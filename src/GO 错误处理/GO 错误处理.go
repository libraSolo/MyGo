package main

import "fmt"

// GO 错误处理
// GO 通过内置的错误接口提供了错误处理机制

// error 类型是一个接口类型
/*
type error interface {
	Error() string
}
*/

// 实例

// 定义一个 DivideError 结构
type DivideError struct {
	etype   int // 错误类型
	dividee int // 除数和被除数
	divider int
}

// 实现 error 接口
func (divideError DivideError) Error() string {
	if 0 == divideError.etype {
		return "除零错误"
	} else {
		return "其他错误"
	}
}

// 定义 除法
func div(a int, b int) (int, *DivideError) {
	if b == 0 {
		return 0, &DivideError{0, a, b}
	} else {
		return a / b, nil
	}
}

func main() {
	// 正确调用
	v, r := div(100, 2)
	if nil != r {
		fmt.Println("(1)fall", r)
	} else {
		fmt.Println("(1)success", v)
	}

	// 错误调用
	v, r = div(100, 0)
	if nil != r {
		fmt.Println("(2)fall", r)
	} else {
		fmt.Println("(2)success", v)
	}
}
