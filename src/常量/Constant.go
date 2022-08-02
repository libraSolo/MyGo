package main

import "fmt"

// 常量
// 常量是不会被修改的量,其类型为:布尔型、数字型（整数型、浮点型和复数）和字符串型

func main() {
	const constA string = "1111"
	const constB = 2
	fmt.Println(constA, constB)

	// 常量可以用函数表达式来计算,不过函数必须是内置函数
	const constC = len(constA)
	fmt.Println(constC)

	// iota 特殊常量, 可以被编译器修改的常量
	// iota 在 const 关键字出现时将被重置为 0 (const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次 (iota 可理解为 const 语句块中的行索引)。
	const (
		a = iota // 0
		b        // 1
		c        // 2
		d = "ha" // 独立值，iota += 1
		e        // "ha"   iota += 1
		f = 100  // iota +=1
		g        // 100  iota +=1
		h = iota // 7, 恢复计数
		i        // 8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
