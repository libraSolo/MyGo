package main

import "fmt"

var c = make(chan int)

func main() {
	go func() {
		for i := 0; i < 2; i++ {
			c <- i
		}
		close(c)
	}()

	r := <-c
	fmt.Println(r)
	r = <-c
	fmt.Println(r)
	r = <-c
	fmt.Println(r) //fatal error: all goroutines are asleep - deadlock! 通道没有关闭 需要写 close(c)

	for v := range c {
		fmt.Println(v) // 没有写的数据 为默认值
	}

}
