/* package main

import "fmt"

func f1() {
	score := 90

	if score >= 60 && score <= 70 {
		fmt.Println("C")
	} else if score > 70 && score <= 90 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}

func f2() {
	if score := 90; score >= 60 && score <= 70 {
		fmt.Println("C")
	} else if score > 70 && score <= 90 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}

//  Monday Tuesday Wednesday Thursday Friday Saturday Sunday

func f3() {
	var c string
	fmt.Println("请输入一个字符：")
	fmt.Scan(&c)
	if c == "M" {
		fmt.Println("Monday")
	} else if c == "T" {
		fmt.Println("请输入第二个字符：")
		fmt.Scan(&c)
		if c == "u" {
			fmt.Println("Tuesday")
		} else {
			fmt.Println("Thursday")
		}
	} else if c == "W" {
		fmt.Println("Wednesday")
	} else if c == "F" {
		fmt.Println("Friday")
	} else if c == "S" {
		fmt.Println("请输入第二个字符：")
		fmt.Scan(&c)
		if c == "a" {
			fmt.Println("Saturday")
		} else {
			fmt.Println("Sunday")
		}
	}
}

func main() {
	// f1()
	// f2()
	f3()
}
*/