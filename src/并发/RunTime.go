package main

import (
	"fmt"
	"runtime"
	"time"
)

func show(msg string) {
	for i := 0; i < 2; i++ {
		fmt.Println(msg)
	}
}

func show2(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		if i > 2 {
			runtime.Goexit() // 退出协程
		}
	}
}

func main() {
	go show("hello") // 子协程来执行

	for i := 0; i < 2; i++ {
		runtime.Gosched() // 让子协程先执行
		fmt.Println("world")
	}
	fmt.Println("end...")

	fmt.Println("--------------------------------")
	go show2("hello2")
	time.Sleep(time.Second) // 等待一秒

	fmt.Println("--------------------------------")
	fmt.Println("CPU 当前核心数", runtime.NumCPU())
	runtime.GOMAXPROCS(2) // 修改核心数
}
