package main

import (
	"fmt"
	"time"
)

// Timer 只执行一次, Ticker 可以周期性执行

func main() {
	ticker := time.NewTicker(time.Second) // 每秒发送当前时间

	count := 1
	for i := range ticker.C {
		fmt.Println("ticker", i)
		count++
		if count >= 5 {
			//ticker.Stop() // 停止 ticker
			break
		}
	}

	fmt.Println("------------------------------------------------")
	chanInt := make(chan int)
	go func() {
		for _ = range ticker.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			}
		}
	}()

	sum := 0
	for v := range chanInt {
		sum += v
		fmt.Printf("v: %v, sum: %v\n", v, sum)

		if sum >= 10 {
			break
		}
	}

}
