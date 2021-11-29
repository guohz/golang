/* package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 获得当前正在运行的进程id
	fmt.Printf("os.Getpid(): %v\n", os.Getpid())
	// 父id
	fmt.Printf("os.Getppid(): %v\n", os.Getppid())

	//设置新进程的属性
	attr := &os.ProcAttr{
		//files指定新进程继承的活动文件对象
		//前三个分别为，标准输入、标准输出、标准错误输出
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		//新进程的环境变量
		Env: os.Environ(),
	}

	//开始一个新进程
	p, err := os.StartProcess("C:\\Windows\\System32\\notepad.exe", []string{"C:\\Windows\\System32\\notepad.exe", "D:\\a.txt"}, attr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	fmt.Println("进程ID：", p.Pid)

	//通过进程ID查找进程
	p2, _ := os.FindProcess(p.Pid)
	fmt.Println(p2)

	//等待10秒，执行函数
	time.AfterFunc(time.Second*10, func() {
		//向p进程发送退出信号
		p.Signal(os.Kill)
	})

	//等待进程p的退出，返回进程状态
	ps, _ := p.Wait()
	fmt.Println(ps.String())
}
*/