package main

import "fmt"

func main() {
	/* 运算符 */
	// 算数运算符
	// 加,减,乘,除,取余
	// 类似 Python 但是比之多了 自增(++)和自减(--)
	a := 10
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)
	/* 关系运算符(比较大小) */
	// < <= >= >
	// 同 Python

	/* 逻辑运算符	*/
	// &&: AND	||: OR	!: NOT

	/* 位运算 */
	// &与, |或, ^异或, <<左移, >>右移,

	/* 其他运算符 */
	// &: 返回储存地址 *: 返回指针变量
	var aa int = 1
	var ptr *int

	ptr = &aa //'ptr' 包含了 'aa' 变量的地址

	fmt.Print(aa, *ptr)
}
