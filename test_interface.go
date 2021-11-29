/* package main

import "fmt"

// Writer Reader

type USB interface {
	read()
	write()
}

type Computer struct {
	name string
}

type Mobile struct {
	model string
}

func (m Mobile) read() {
	fmt.Printf("m.model: %v\n", m.model)
	fmt.Println("read...")
}

func (m Mobile) write() {
	fmt.Printf("m.model: %v\n", m.model)
	fmt.Println("write...")
}

func (c Computer) read() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("read...")
}

func (c Computer) write() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("write...")
}

func main() {
	c := Computer{
		name: "联想",
	}

	c.read()
	c.write()

	m := Mobile{
		model: "5G",
	}

	m.read()
	m.write()
}
*/