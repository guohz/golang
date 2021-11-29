/* package main

import "fmt"

func cal(base int) (func(int) int, func(int) int) {
	add := func(a int) int {
		base += a
		return base
	}

	sub := func(a int) int {
		base -= a
		return base
	}
	return add, sub

}

func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func main() {

	add, sub := cal(100)

	r := add(100)
	fmt.Printf("r: %v\n", r)
	r = sub(50)
	fmt.Printf("r: %v\n", r)

	add1, sub1 := cal(100)

	r = add1(1)
	fmt.Printf("r: %v\n", r)
	r = sub1(2)
	fmt.Printf("r: %v\n", r)

	/* f := add()
	r := f(10)
	fmt.Printf("r: %v\n", r)
	r = f(20)
	fmt.Printf("r: %v\n", r)
	r = f(30)
	fmt.Printf("r: %v\n", r)
	fmt.Println("------------")

	f1 := add()
	r = f1(100)
	fmt.Printf("r: %v\n", r)
	r = f1(100)
	fmt.Printf("r: %v\n", r) */
}
 */