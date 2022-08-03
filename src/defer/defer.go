package main

import "fmt"

// defer 将后面的语句 延迟处理
// defer 延迟处理后的语句 按照逆序处理

// defer 用途
// 关闭文件, 锁资源释放, 数据库连接释放

func f1() {
	fmt.Println("1...")
	defer fmt.Println("2...")
	defer fmt.Println("3...")
	fmt.Println("4...")
}

func main() {
	f1() // 1 4 3 2
}
