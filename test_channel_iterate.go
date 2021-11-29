/* package main

import "fmt"

var c = make(chan int)

func main() {

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		v, ok := <-c
		if ok {
			fmt.Printf("v: %v\n", v)
		} else {
			break
		}
	}

	/* 	for v := range c {
		fmt.Printf("v: %v\n", v)
	} */

	/* for i := 0; i < 3; i++ {
		r := <-c
		fmt.Printf("r: %v\n", r)
	} */

}
 */