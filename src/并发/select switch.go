package main

import (
	"fmt"
	"time"
)

// select switch 仅用于处理异步IO操作
// select 的 case 语句必须是一个 channel 操作
// select 的 default 子句总是可运行的
// 多个 case 可执行, 随机执行
// 没有case 且 没有 default , select 将阻塞,直到某个 case 可执行

var chanInt = make(chan int)
var chanStr = make(chan string)

func main() {
	go func() {
		chanInt <- 100
		chanStr <- "Hello"

		close(chanInt)
		close(chanStr)
	}()

	for {
		select {
		case r := <-chanInt:
			fmt.Printf("chanInt: %v\n", r)
		case r := <-chanStr:
			fmt.Printf("chanStr: %v\n", r)
		default:
			fmt.Println("default")
		}
		time.Sleep(time.Second)
	}
}
