/* package main

import "fmt"

type Pet interface {
	eat()
	sleep()
}

type Dog struct {
}

type Cat struct {
}

// Dog 实现Pet接口

func (dog Dog) eat() {
	fmt.Println("dog eat...")
}
func (dog Dog) sleep() {
	fmt.Println("dog sleep...")
}

// Cat 实现Pet接口
func (cat Cat) eat() {
	fmt.Println("cat eat...")
}
func (cat Cat) sleep() {
	fmt.Println("cat sleep...")
}

type Person struct {
}

// pet 既可以传递Dog也可以传递Cat
func (person Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func main() {
	dog := Dog{}
	cat := Cat{}

	person := Person{}
	person.care(dog)
	person.care(cat)
}
*/