/* package main

import (
	"fmt"
	"os"
)

func test1() {
	var name string
	fmt.Println("请输入姓名：")
	fmt.Scan(&name)
	fmt.Printf("name: %v\n", name)
}

func test2() {
	const name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())
}

func test3() {
	const name, age = "Kim", 22
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")

	// The n and err return values from Fprint are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")
}

func main() {
	// test1()
	// test2()
	test3()
}
*/