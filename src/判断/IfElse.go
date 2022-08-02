package main

import "fmt"

func main() {
	var a = 16
	var b string

	// if else
	// 每一个 if else 都需要加入括号 同时 else 位置不能在新一行
	if a == 20 {
		fmt.Println("high")
	} else if a < 20 && a > 15 { // 与 Python 不同 此处用的 &&
		fmt.Println("mid")
	} else {
		fmt.Println("low")
	}

	// Go 中没有 While 语句
	// switch
	switch a {
	case 20:
		b = "high"
	case 16:
		b = "mid"
	case 10:
		b = "low"
	default: // 不论放在那里,都是最后执行
		b = "None"
	}

	switch {
	case b == "high":
		fmt.Println("high")
	case b == "mid":
		fmt.Println("mid")
	case b == "low":
		fmt.Println("low")
	}

	// 判断类型
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Println("x 的类型 :", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func (int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型")
	default:
		fmt.Println("未知型")
	}

	// fallthrough 强制执行后面的 case 语句
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}

	// select
	// 每个 case 都必须是一个通信
	// 类似 switch ,select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("received", i1, "from c1")
	case c2 <- i2:
		fmt.Println("sent", i2, "to c2")
	case i3, ok := (<-c3):
		if ok {
			fmt.Println("received", i3, "from c3")
		} else {
			fmt.Println("c3 is closed")
		}
	default:
		fmt.Println("no communication")
	}
}
