/* package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
}

func Marshal() {
	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}
	// b, _ := xml.Marshal(p)
	// 有缩进格式
	b, _ := xml.MarshalIndent(p, " ", "  ")
	fmt.Printf("%v\n", string(b))
}
func Unmarshal() {
	b1 := []byte(`<person><name>tom</name><age>20</age><email>tom@gmail.com</email></person>`)
	var m Person
	xml.Unmarshal(b1, &m)
	fmt.Printf("m: %v\n", m)
}

func read() {
	/*
		<?xml version="1.0" encoding="UTF-8"?>
		<person>
		   <name>tom</name>
		   <age>20</age>
		   <email>tom@gmail.com</email>
		</person>
	*/
	b, _ := ioutil.ReadFile("a.xml")
	var p Person
	xml.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p)
}

func write() {
	type Person struct {
		XMLName xml.Name `xml:"person"`
		Name    string   `xml:"name"`
		Age     int      `xml:"age"`
		Email   string   `xml:"email"`
	}

	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}

	f, _ := os.OpenFile("a.xml", os.O_WRONLY, 0777)
	defer f.Close()
	e := xml.NewEncoder(f)
	e.Encode(p)
}

func main() {
	// Marshal()
	// Unmarshal()
	// read()
	write()
}
 */