package main

import (
	"fmt"
	"sync"
)

var wp sync.WaitGroup

func showMsg(s int) {
	defer wp.Add(-1)
	println(s)
}

func main() {

	for i := 0; i < 10; i++ {
		// 启动一个协程
		go showMsg(i)
		wp.Add(1)
	}
	wp.Wait() // 等待为 0
	fmt.Println("end...")
}
