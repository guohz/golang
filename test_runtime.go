/* package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("a:", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("b:", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	go a()
	go b()

	time.Sleep(time.Second)

}

/* package main

import (
	"fmt"
	"runtime"
	"time"
)

func show() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			runtime.Goexit()
		}
	}
}

func main() {
	go show()
	time.Sleep(time.Second)
}

*/

/* package main

import (
	"fmt"
	"runtime"
)

func show(msg string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("msg: %v\n", msg)
	}
}

func main() {
	go show("java") // 子协程来运行
	for i := 0; i < 2; i++ {
		runtime.Gosched() // 我有权利执行任务了，让给你（其他子协程来执行）
		fmt.Printf("\"golang\": %v\n", "golang")
	}
	fmt.Println("end...")
}
*/
 */