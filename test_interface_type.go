/* package main

import "fmt"

type Player interface {
	playMusic()
}

type Video interface {
	playVideo()
}

type Mobile struct {
}

func (m Mobile) playMusic() {
	fmt.Println("play music")
}

func (m Mobile) playVideo() {
	fmt.Println("play video")
}

type Pet interface {
	eat()
}

type Dog struct {
}

type Cat struct {
}

func (dog Dog) eat() {
	fmt.Println("dog eat...")
}

func (cat Cat) eat() {
	fmt.Println("cat eat...")
}

func main() {
	m := Mobile{}
	m.playMusic()
	m.playVideo()

	/* dog := Dog{}
	cat := Cat{}
	dog.eat()
	cat.eat() */
	var pet Pet
	pet = Dog{}
	pet.eat()
	pet = Cat{}
	pet.eat()
}
 */