package main

import (
	"bytes"
	"fmt"
	"strings"
)

type WebSite struct {
	Name string
}

func main() {

	/* 字符串的声明 */
	var s string = "hello world"
	var s1 = "hello world"
	s2 := "hello world"
	s3 := `
		hello world
		`

	/* 字符串连接 */
	s1 := "hello"
	s2 := "world"
	fmt.Printf("%s,%s\n", s1, s2) // hello,world

	s3 := strings.Join([]string{s1, s2}, ",")
	fmt.Println(s3) // hello,world

	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString(",")
	buffer.WriteString(s2)
	fmt.Println(buffer.String()) //hello,world

	/* 转义字符 */
	// \n:换行 \t:tab 4个空格  \r 回车(返回行首)

	/* 切片(索引) */
	// 切片.go

	/* 字符串函数 */
	// strings.len(string):求长度
	// strings.Split(string,s): 分割
	fmt.Println(strings.Split("hello,world", ","))
	// strings.Contains(string,s): 包含字符串
	fmt.Println(strings.Contains("hello,world", "h"))
	// strings.ToLower(s): 全小写
	// strings.ToUpper(s): 全大写
	// strings.HasPrefix(string,s): 以 s 开头
	// strings.HasSuffix(string,s): 以 s 结尾
	// strings.Index(string,s): 查找 s 的索引位置

	/* 格式化输出 */
	site := WebSite{"mucun"}
	// %v: 变量
	fmt.Printf("%v \n", site)
	// %T: 类型
	fmt.Printf("%T \n", site)
	// %#v: 相应值的 Go 语法表示
	fmt.Printf("%#v \n", site)
	// %%: % 输出
	// %t: bool类型输出
	// %c: Unicode码
	// %b: 二进制表示
}
