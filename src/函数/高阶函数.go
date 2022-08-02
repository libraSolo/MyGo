package main

import "fmt"

func sayHello(s string) {
	fmt.Println(s)
}

func test(name string, f func(string)) {
	f(name)
}

// ----------------------------------------------------------------
func add(a, b int) int {
	return a + b
}

func cal() func(int, int) int {
	return add
}

func main() {
	test("hello", sayHello)

	num := cal()
	r := num(1, 2)
	fmt.Println(r)
}
