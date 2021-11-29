/* package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var i int32 = 100

// cas compare and swap  old new

func add() {
	atomic.AddInt32(&i, 1)
}

func sub() {
	atomic.AddInt32(&i, -1)
}

/* var i = 100

// 100 万 你 -10 你老婆 100-10 = 90

var lock sync.Mutex

func add() {
	lock.Lock()
	i++
	lock.Unlock()
}

func sub() {
	lock.Lock()
	i--
	lock.Unlock()
} */

func main() {

	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}

	time.Sleep(time.Second * 2)

	fmt.Printf("i: %v\n", i)

}
 */