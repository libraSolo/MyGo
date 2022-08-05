package main

import (
	"fmt"
	"time"
)

// Timer 定时器,实现定时操作

func main() {
	timer := time.NewTimer(time.Second * 2)
	fmt.Printf("%v\n", time.Now())
	t1 := <-timer.C // timer 是阻塞的,直到时间到了
	// <- timer.C // 也是等待...
	fmt.Printf("%v\n", t1)

	time.Sleep(time.Second * 2) // 等待 2s
	fmt.Println(time.Now())

	<-time.After(time.Second * 2) // 返回值是 chan Time
	fmt.Println(time.Now())

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println(time.Now())
	}()
	stop := timer2.Stop() // 停止定时器
	// 阻止 timer 事件发生, timer 计时器停止,相应的事件也不执行
	if stop {
		fmt.Println("stop")
	}

	fmt.Println(time.Now())
	timer3 := time.NewTimer(time.Second * 5)
	timer3.Reset(time.Second) // 重设时间
	<-timer3.C
	fmt.Println("timer3", time.Now())

	time.Sleep(time.Second * 2)
	fmt.Println("main end...")
}
