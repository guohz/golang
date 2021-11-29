/* package main

import "fmt"

func test2() {
	var s1 = []int{1, 2, 3, 4, 5}

	/* for i, v := range s1 {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	} */

	for _, v := range s1 {
		fmt.Printf("v: %v\n", v)
	}
}

func test1() {
	var s1 = []int{1, 2, 3, 4, 5}
	// l := len(s1)
	/* 	for i := 0; i < l; i++ {
		fmt.Printf("s1[%v]: %v\n", i, s1[i])
	} */
	for i := 0; i < len(s1); i++ {
		fmt.Printf("s1[%v]: %v\n", i, s1[i])
	}
}
func main() {
	// test1()
	test2()
}
 */