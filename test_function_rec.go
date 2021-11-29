/* package main

import "fmt"

func f2(a int) int {
	if a == 1 {
		// 1. 退出条件
		return 1
	} else {
		// 2. 自己调用自己
		return a * f2(a-1)
	}
}

func f1() {
	var s int = 1
	//  5! = 5x4x3x2x1
	for i := 1; i <= 5; i++ {
		s *= i
	}
	fmt.Printf("s: %v\n", s)
}

func main() {
	// f1()
	i := f2(6)
	fmt.Printf("i: %v\n", i)
}
*/