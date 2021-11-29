/* package main

import "fmt"

func test1() {
	var name string
	name = "tom"
	var p_name *string
	p_name = &name

	fmt.Printf("name: %v\n", name)
	fmt.Printf("p_name: %v\n", p_name)
	fmt.Printf("p_name: %v\n", *p_name)
}

func test2() {
	type Person struct {
		id   int
		name string
		age  int
	}

	tom := Person{
		id:   101,
		name: "tom",
		age:  20,
	}

	var p_person *Person
	p_person = &tom

	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("p_person: %p\n", p_person)
	fmt.Printf("p_person: %v\n", *p_person)
}

func test3() {
	type Person struct {
		id   int
		name string
		age  int
	}

	var tom = new(Person)
	/* (*tom).id = 101
	(*tom).name = "tom" */
	tom.id = 101
	tom.name = "tom"

	fmt.Printf("tom: %v\n", *tom)
}

func main() {
	// test2()
	test3()
}
 */