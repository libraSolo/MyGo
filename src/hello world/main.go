package main // 定义包名
import (
	"fmt"
) // 导包

// GO 的数据类型
/*
	布尔: true, false
	数字: int, uint8, float32, float64, 复数, ...
	字符串: string
	派生类: 指针类型, 数组类型, 结构化类型, ...
*/
// 声明变量并初始化
var a string = "mucun"

// 声明多个变量并初始化
var b, c int = 1, 2

// 声明变量, 没有初始化,默认值为 0
var d int    // 0
var e string // ''
var f bool   // false

// 声明变量, 自行判断变量类型
var g = 1

// 对于声明过的语句再次声明,则会报错
// g := 2	此时会报错 ’:=‘ 相等于 var g int  = 1

// 因式分解关键字写法一般用于 全局变量
var (
	globalA string
	globalB int
)

func main() {
	fmt.Println("hello world")

}
