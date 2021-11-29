/* package main

import (
	"fmt"
	"sync"
)

// import "sync"

var wp sync.WaitGroup

func showMsg(i int) {
	// defer wp.Add(-1)
	defer wp.Done()
	fmt.Printf("i: %v\n", i)
}

func main() {

	for i := 0; i < 10; i++ {
		// 启动一个协程来执行
		go showMsg(i)
		wp.Add(1)
	}

	wp.Wait()
	// 主协程
	fmt.Println("end...")

}
*/