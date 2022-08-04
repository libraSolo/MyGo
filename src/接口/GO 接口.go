package main

import "fmt"

// 定义接口
type Phone interface {
	call()
	call2()
}

// 定义结构体
type Phone1 struct {
	id           int
	name         string
	categoryID   int
	categoryName string
}

// 第一个类的第一个回调
func (test Phone1) call() {
	fmt.Println("test Phone1 call", Phone1{id: 1, name: "诺基亚"})
}

// 第一个类的第二个回调
func (test Phone1) call2() {
	fmt.Println("test Phone1 call2", Phone1{id: 2, name: "苹果"})
}

// 定义结构体二
type Phone2 struct {
	memberID   int
	memberName string
}

// 第二个类的第一个回调
func (test2 Phone2) call() {
	fmt.Println("test2 Phone1 call", Phone2{memberID: 1, memberName: "小明"})
}

// 第二个类的第二个回调
func (test2 Phone2) call2() {
	fmt.Println("test2 Phone1 call2", Phone2{memberID: 2, memberName: "小红"})
}

func main() {
	var phone Phone

	// 实例化第一个接口
	phone = new(Phone1)
	phone.call()  // test Phone1 call {1 诺基亚 0 }
	phone.call2() // test Phone1 call2 {2 苹果 0 }

	// 实例化第二个接口
	phone = new(Phone2)
	phone.call()  // test2 Phone1 call {1 小明}
	phone.call2() // test2 Phone1 call2 {2 小红}
}

//
