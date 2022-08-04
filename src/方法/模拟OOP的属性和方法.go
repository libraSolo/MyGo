package main

import "fmt"

type Person2 struct {
	name string
	age  int
}

func (per Person2) eat() {
	fmt.Println("eat...")
}

func main() {
	per := Person2{"name", 20}
	println(per.name)
}
