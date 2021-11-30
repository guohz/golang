/* package main

import "fmt"

func testAppend() {
	s := []int{1, 2, 3}
	i := append(s, 100)
	fmt.Printf("i: %v\n", i)
	s1 := []int{4, 5, 6}
	i2 := append(s, s1...)
	fmt.Printf("i2: %v\n", i2)
}

func testLen() {
	s := "Hello world"
	fmt.Printf("len(s): %v\n", len(s))
	s1 := []int{1, 2, 3}
	fmt.Printf("len(s1): %v\n", len(s1))
}

func testPrint() {
	name := "tom"
	age := 20
	print(name, " ", age, "\n")
	fmt.Println("------------")
	println(name, " ", age)
}

func testPanic() {
	defer fmt.Println("panic 后 我还会执行...")
	panic("哎呀 完蛋了...")
	fmt.Println("end...")
}

func testNew() {
	b := new(bool)
	fmt.Printf("b: %T\n", b)
	fmt.Printf("b: %v\n", *b)
	i := new(int)
	fmt.Printf("i: %T\n", i)
	fmt.Printf("i: %v\n", *i)

	s := new(string)
	fmt.Printf("s: %T\n", s)
	fmt.Printf("s: %v\n", *s)
}

func testMakeVsNew() {
	var p *[]int = new([]int)
	fmt.Printf("p: %v\n", p)
	v := make([]int, 10)
	fmt.Printf("v: %v\n", v)
}

func main() {

}
*/