package main

import "fmt"

func main() {
	type Dog struct {
		name string
		age  int
	}

	type Person struct {
		dog  Dog
		name string
	}

	dog := Dog{"dog", 2}
	per := Person{dog, "小明"}

	fmt.Println(per) // {{dog 2} 小明}

}
