package main

import (
	"fmt"
	"sync/atomic"
)

// atomic 提供了原子操作 保证了任一时刻止只有一个 goroutines 对变量进行操作
var i2 int32 = 100

func add2() {
	atomic.AddInt32(&i2, 1) // 乐观锁(当前值和过去值比较,一样再进行加减)
}
func sub2() {
	atomic.AddInt32(&i2, -1)
}

func main() {
	for i := 0; i < 100; i++ {
		add2()
		sub2()
	}
	atomic.LoadInt32(&i2) // 以原子方式读
	fmt.Println(i2)
	atomic.StoreInt32(&i2, 1) // 以原子方式写
	fmt.Println(i2)
}
