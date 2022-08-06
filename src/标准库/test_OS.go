package main

import (
	"fmt"
	"os"
)

// 创建文件
func test_create() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(f)
	}
}

// 创建目录
func test_dir() {
	err := os.Mkdir("test", os.ModePerm)
	if err != nil {
		fmt.Println("error: ", err)
	}
}

// 删除
func test_remove() {
	err := os.Remove("test.txt")
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func main() {
	os.Getwd()                                             // 获取当前目录
	os.Chdir("D:/")                                        // 改变当前目录
	os.TempDir()                                           // 获取当前临时目录
	os.Rename("test.txt", "test2.txt")                     // 重命名
	file, err := os.ReadFile("test.txt")                   // 读取文件
	os.WriteFile("test.txt", []byte("hello"), os.ModePerm) // 写入类型:字节数组

}
