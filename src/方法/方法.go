package main

import "fmt"

// Go 方法是一种特殊的函数, 定义于 struct 之上(与 struct 关联,绑定)
/*		语法格式
type mytype struct {}
func (recv mytype) my_method() return_type {}
*/

type Person struct {
	name string
}

func (per Person) eat() {
	fmt.Printf("%v,eat...", per.name)
}

func main() {
	per := Person{"DaMing"}
	per.eat()
}
