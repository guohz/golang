/* package main

import "fmt"

type Person struct {
	id   int
	name string
}

func showPerson2(per *Person) {
	per.id = 102
	per.name = "kite"
	// fmt.Printf("per: %v\n", per)
}

// 值传递拷贝了一份副本
func showPerson(per Person) {
	per.id = 101
	per.name = "kite"
	fmt.Printf("per: %v\n", per)
}

func main() {
	tom := Person{
		id:   100,
		name: "tom",
	}

	per := &tom

	fmt.Printf("tom: %v\n", tom)

	showPerson2(per)

	fmt.Println("-----------")
	fmt.Printf("per: %v\n", *per)

	/* fmt.Printf("tom: %v\n", tom)
	fmt.Println("-------------")
	showPerson(tom)
	fmt.Printf("tom: %v\n", tom) */
}
 */