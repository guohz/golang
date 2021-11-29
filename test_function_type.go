/* package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	type f1 func(int, int) int
	var ff f1
	ff = sum
	r := ff(1, 2)
	fmt.Printf("r: %v\n", r)

	ff = max
	r = ff(1, 2)
	fmt.Printf("r: %v\n", r)
}
*/