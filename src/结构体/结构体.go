package main

import "fmt"

/* 结构体 */
/*
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
*/

type Books struct {
	title  string
	author string
	bookID int
}

func main() {
	// 创建一个新的结构体
	fmt.Println(Books{"golang", "mucun", 10001})
	// 访问结构体成员
	// 结构体.成员名

	var book1 Books
	book1.title = "GO"
	book1.author = "mucun"
	book1.bookID = 10001

	fmt.Println(book1)
}
