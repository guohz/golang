/* package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func test1() {
	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}

	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b))
}

func test2() {
	b := []byte(`{"Name":"tom","Age":20,"Email":"tom@gmail.com"}`)
	var p Person
	json.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p)
}

// 解析嵌套类型
func test3() {
	b := []byte(`{"Name":"tom","Age":20,"Email":"tom@gmail.com", "Parents":["big tom", "kite"]}`)
	// var f interface{}
	var f map[string]interface{}
	json.Unmarshal(b, &f)
	fmt.Printf("f: %v\n", f)

	for k, v := range f {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}
}

func test4() {
	type Person struct {
		Name   string
		Age    int
		Email  string
		Parent []string
	}

	p := Person{
		Name:   "tom",
		Age:    20,
		Email:  "tom@gmail.com",
		Parent: []string{"big tom", "big kite"},
	}

	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b))
}

func test6() {
	f, _ := os.Open("a.json")
	defer f.Close()
	d := json.NewDecoder(f)
	var v map[string]interface{}
	d.Decode(&v)

	fmt.Printf("v: %v\n", v)
}

func test7() {
	type Person struct {
		Name   string
		Age    int
		Email  string
		Parent []string
	}

	p := Person{
		Name:   "tom",
		Age:    20,
		Email:  "tom@gmail.com",
		Parent: []string{"big tom", "big kite"},
	}

	f, _ := os.OpenFile("a.json", os.O_WRONLY, 0777)
	defer f.Close()
	e := json.NewEncoder(f)
	e.Encode(p)

}

func main() {
	// test2()
	// test3()
	// test4()
	// test6()
	// test3()
	test7()
}
*/