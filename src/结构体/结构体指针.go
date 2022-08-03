package main

import "fmt"

type Tom struct {
	name string
	id   int
	age  int
}

func main() {
	tom := Tom{"tom", 1, 20}
	var p_tom *Tom
	p_tom = &tom

	fmt.Printf("%p\n", p_tom)  // 0xc0000223e0
	fmt.Printf("%v\n", *p_tom) // {tom 1 20}

	// 可以使用 new 创建结构体指针
	var p_tom2 = new(Tom)
	fmt.Printf("%T\n", p_tom2)  // *main.Tom
	fmt.Printf("%v\n", *p_tom2) // { 0 0}
}
