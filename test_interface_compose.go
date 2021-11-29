/* package main

import "fmt"

type Flyer interface {
	fly()
}

type Swimmer interface {
	swim()
}

// 接口的组合
type FlyFish interface {
	Flyer
	Swimmer
}

type Fish struct {
}

func (fish Fish) fly() {
	fmt.Println("fly")
}

func (fish Fish) swim() {
	fmt.Println("swim...")
}

func main() {

	var ff FlyFish
	ff = Fish{}
	ff.fly()
	ff.swim()
}
*/