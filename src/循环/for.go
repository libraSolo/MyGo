package main

import "fmt"

func main() {
	// 计算 1到100的数字和
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 此写法类似 while 语句
	sum1 := 1
	for sum1 <= 10 { //for; sum1 <= 10; 此写法有分号
		sum1 += sum1
	}
	fmt.Println(sum1)

	// 无限循环
	//sum2 := 0
	//for {
	//	sum2++
	//}

	// For-each range 循环
	strings := []string{"hello", "world"}
	for i, j := range strings {
		fmt.Println(i, j)
	}

	numbers := [6]int{1, 2, 3, 4, 5}
	for i, j := range numbers {
		fmt.Printf("第 %d ：%d \n", i, j)
	}

	// for 循环读取 key or value
	// 读取 key 和 value
	for key, value := range numbers {
		fmt.Println(key, value)
	}

	// 读取 key
	for key := range numbers {
		fmt.Println(key)
	}

	// 读取 value
	for _, value := range numbers {
		fmt.Println(value)
	}
}
