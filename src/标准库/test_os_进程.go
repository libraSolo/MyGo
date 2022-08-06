package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("当前进程id ", os.Getpid())
	fmt.Println("父id ", os.Getppid())

	// 设置新进程的属性
	attr := &os.ProcAttr{
		// files 指定新进程继承的活动文件对象
		// 前三个为, 标准输出,标准输入,标准错误输出
		Files: []*os.File{os.stdin, os.stdout, os.Stderr},
		// 新进程环境变量
		Env: os.Environ(),
	}

	// 开始一个新进程
	p, err := os.StartProcess("E:\\")
	// 通过进程ID 查找进程
	p2, _ := os.FindProcess(p.Pid)
	// 向进程P发送退出信号
	p.Signal(os.Kill)
	// 等待进程p的退出,过程阻塞
	ps, _ := p.Wait()
}
