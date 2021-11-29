/* package main

import "fmt"

func test5() {
	var a1 = [...]int{1, 2, 3}
	fmt.Printf("a1: %v\n", a1)
	a1[0] = 100
	fmt.Printf("a1: %v\n", a1)
}

func test4() {
	var a1 = [...]int{0: 1, 3: 100, 5: 10}
	var a2 = [...]bool{2: true, 5: false}
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
}

func test3() {
	// 数组的初始化，默认长度省略长度 ...
	var a1 = [...]int{1, 2, 3}
	fmt.Printf("len(a1): %v\n", len(a1))
	var a2 = [...]string{"duoke360.com", "golang"}
	fmt.Printf("a2: %v\n", a2)
}

func test2() {
	// 数组的初始化
	// 初始化列表
	var a1 = [3]int{1, 2}
	fmt.Printf("a1: %v\n", a1)
	var a2 = [2]string{"hello", "world"}
	fmt.Printf("a2: %v\n", a2)
	var a3 = [2]bool{true, false}
	fmt.Printf("a3: %v\n", a3)
}

func test1() {
	var a1 [2]int
	var a2 [3]string
	var a3 [2]bool

	fmt.Printf("a1: %T\n", a1)
	fmt.Printf("a2: %T\n", a2)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	test5()
}
*/