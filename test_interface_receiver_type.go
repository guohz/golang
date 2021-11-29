/* package main

import "fmt"

type Pet interface {
	eat(string) string
}

type Dog struct {
	name string
}

/* func (dog Dog) eat(name string) string {
	dog.name = "花花..."
	fmt.Printf("name: %v\n", name)
	return "吃的很好"
} */

func (dog *Dog) eat(name string) string {
	dog.name = "花花..."
	fmt.Printf("name: %v\n", name)
	return "吃的很好"
}

func main() {

	/* 	dog := Dog{
	   		name: "花花",
	   	}

	   	s := dog.eat("火鸡")
	   	fmt.Printf("s: %v\n", s)
	   	fmt.Printf("dog: %v\n", dog)
	*/

	dog := &Dog{
		name: "花花",
	}

	s := dog.eat("火鸡")
	fmt.Printf("s: %v\n", s)
	fmt.Printf("dog: %v\n", dog)

}
 */