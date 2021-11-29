/* package main

import "fmt"

func test1() {
	fmt.Println("我既没有参数，也没有返回值...")
}

func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

func sum2(a int, b int) int {
	// ret := a + b
	return a + b
}

func test3() (string, int) {
	/* name = "tom"
	age = 20 */
	n := "tom"
	a := 20
	// return name, age
	return n, a
}

func main() {
	// test1()
	name, age := test3()
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
}
 */