/* package main

import "fmt"

// 形参
func sum(a int, b int) int {
	return a + b
}

// copy
func f1(a int) {
	a = 100
}

func f2(s []int) {
	s[0] = 1000
}

func f3(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

func f4(name string, ok bool, args ...int) {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("ok: %v\n", ok)
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {

	f4("tom", true, 1, 2, 3)

	/* 	f3(1, 2, 3, 4)
	   	f3(3, 4, 5, 6, 7) */
	/* s := []int{1, 2, 3}
	f2(s)
	fmt.Printf("s: %v\n", s) */
	/* x := 200
	f1(x)
	fmt.Printf("x: %v\n", x) */
	/* r := sum(1, 2)
	fmt.Printf("r: %v\n", r) */
}
 */