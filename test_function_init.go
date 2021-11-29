/* package main

import "fmt"

var i int = initVar()

var j int

/* func init() {
	fmt.Println("init2...")
}

func init() {
	fmt.Println("init...")
} */

func init() {
	fmt.Println("init...")
	j = 100
}

func initVar() int {
	fmt.Println("initVar...")
	return 100
}

func main() {
	// init()
	fmt.Println("main....")
}
 */