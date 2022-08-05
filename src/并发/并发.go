package main

import (
	"fmt"
	"time"
)

// GO 支持并发 关键字 goroutine 开启
// goroutine 语法格式
// go 函数名(参数列表)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world") // go启动了一条协程
	say("hello")

	/*
		go say("协程1")
		go say("协程2")
		fmt.Println("协程3")  //主协程,   当主协程结束, 其他携程结束
	*/
}
