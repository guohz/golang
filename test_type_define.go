/* package main

import "fmt"

func main() {
	/* type MyInt int
	var i MyInt
	i = 100

	i: main.MyInt,100
	fmt.Printf("i: %T,%v\n", i, i) */

	type MyInt = int
	var i MyInt
	i = 100
	fmt.Printf("i: %T, %v\n", i, i)
	// i: int, 100
}
 */