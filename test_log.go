/* package main

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

func test1() {
	log.Print("my log")
	log.Printf("my log %d", 100)
	name := "tom"
	age := 20
	log.Println(name, " ", age)
}

func test2() {
	defer fmt.Println("panic 结束后再执行。。。")
	log.Print("my log")
	log.Panic("my panic")
	fmt.Println("end...")
}

func test3() {
	defer fmt.Println("defer...")
	log.Print("my log")
	log.Fatal("fatal...")
	// os.Exit(1)
	fmt.Println("end...")
}

func init() {
	/* log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("MyLog: ")
	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	// defer f.Close()
	if err != nil {
		log.Fatal("日志文件错误")
	}
	log.SetOutput(f) */

	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	// defer f.Close()
	if err != nil {
		log.Fatal("日志文件错误")
	}
	logger = log.New(f, "MyLog: ", log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {
	/* i := log.Flags()
	fmt.Printf("i: %v\n", i)
	log.Print("my log...") */
	logger.Print("my log...")
}
 */