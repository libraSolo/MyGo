package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// 读取目录
func testReadDir() {
	fi, _ := ioutil.ReadDir(".")
	for _, v := range fi {
		fmt.Println(v.Name())
	}
}

// 读文件
func testReadFile() {
	b, _ := ioutil.ReadFile("go.mod")
	fmt.Println(string(b))
}

// 写文件
func testWriteFile() {
	ioutil.WriteFile("test.txt", []byte("hello world"), 0664)
}
func main() {
	r := strings.NewReader("hello world")
	b, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println("Error reading")
	}
	fmt.Printf("%v", string(b))

	testReadDir()
	testReadFile()
	testWriteFile()
}
