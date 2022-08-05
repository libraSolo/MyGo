package main

import "fmt"

// channel
// 通道 是用来传数据的一个数据结构
// 通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯
// 操作符 <- 用于指定通道的方向,发送或接收,如果未指明方向,则为双向通道

/*
ch <- v		// 把 v 发送到通道 ch
v := <- ch	// 从 ch 接收数据 并把值赋给 v
*/

// 声明通道
//ch := make(chan int)		 // 整形无缓存通道
//ch := make(chan int, 10)		// 整形有缓存通道

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	s := []int{1, 2, 3, 4}

	c := make(chan int)
	go sum(s[:len(s)/2], c) // 3
	go sum(s[len(s)/2:], c) // 7
	x, y := <-c, <-c        // 从通道c 中接收
	fmt.Println(x, y, x+y)  // 3 7 10

	// 通道缓冲区
	// 通道可以设置缓冲区, 通过 make 的第二个参数指定缓冲区大小
	// ch := make(chan int, 100)

	ch := make(chan int, 2)
	// 因为 ch 带缓冲性质,可以同时发送两个数据,不用立即需要去同步读取数据
	ch <- 1
	ch <- 2
	// 获取这两个数据
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2

	fmt.Println("--------------------------------")
	// go 遍历通道和关闭通道
	// 如果通道接收不到数据后则接收的数据就为 false，这时通道就可以使用 close() 函数来关闭

	c1 := make(chan int, 10)
	go fibonacci(cap(c1), c1)
	//range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c1 {
		fmt.Println(i)
	}

}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
