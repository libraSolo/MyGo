package main

import (
	"fmt"
	"sync"
)

var i int = 100
var lock sync.Mutex

func add() {
	lock.Lock()
	i += 1
	fmt.Println(i)
	lock.Unlock()
}

func sub() {
	lock.Lock()
	i -= 1
	fmt.Println(i)
	lock.Unlock()
}
func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}
	fmt.Println(i)
}
