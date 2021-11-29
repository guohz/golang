/* package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 创建int类型通道，只能传入int类型值
var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send: %v\n", value)
	time.Sleep(time.Second * 5)
	values <- value
}

func main() {
	// 从通道接收值
	defer close(values)
	go send()
	fmt.Println("wait...")
	value := <-values
	fmt.Printf("receive: %v\n", value)
	fmt.Println("end...")
}
 */