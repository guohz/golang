/* package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println("eat...")
}

func (a Animal) sleep() {
	fmt.Println("sleep...")
}

type Dog struct {
	Animal // 可以理解为继承
	color  string
}

type Cat struct {
	Animal
	bbb string
}

func main() {
	dog := Dog{
		Animal{name: "花花", age: 2},
		"褐色",
	}

	dog.eat()
	dog.sleep()
	fmt.Printf("dog.color: %v\n", dog.color)
	fmt.Printf("dog.age: %v\n", dog.age)

	cat := Cat{
		Animal{name: "嘿嘿", age: 3},
		"bbb",
	}

	cat.eat()
	cat.sleep()

	fmt.Printf("cat: %v\n", cat)
}
*/