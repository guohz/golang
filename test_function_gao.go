/* package main

import "fmt"

func sayHello(name string) {
	fmt.Printf("Hello,%s", name)
}

func test(name string, f func(string)) {
	f(name)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func cal(operator string) func(int, int) int {
	switch operator {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	// test("tom", sayHello)

	ff := cal("+")
	r := ff(1, 2)
	fmt.Printf("r: %v\n", r)
	ff = cal("-")
	r = ff(2, 1)
	fmt.Printf("r: %v\n", r)
}
*/