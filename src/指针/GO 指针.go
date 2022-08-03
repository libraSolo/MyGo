package main

import "fmt"

// 指针指向了一个值的内存地址
// & 取地址		* 根据地址取值
// 指针的声明如下
// var var_name *var_type

var ip *int     // 指向整形
var fp *float64 // 指向浮点型

// 指针类型前加上 *(前缀) 可获取指针指向的内容

func main() {
	a := 10
	ip = &a

	// GO 的取地址符是 &,返回变量的内存地址
	fmt.Println(&a)  // 0xc000018098
	fmt.Println(ip)  // 0xc000018098
	fmt.Println(*ip) // 10

	// 空指针:当一个指针被定义后没有分配到任何变量,其值为 nil
	// nil 指针被成为 空指针,指代零值或空值
	// 一个指针变量通常缩写为 ptr
	fmt.Println(fp) // <nil>
}
