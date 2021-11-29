/* package main

import (
	"fmt"
	"os"
)

// 创建文件
func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

// 创建目录
func makeDir() {
	/* err := os.Mkdir("a", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} */

	err2 := os.MkdirAll("a/b/c", os.ModePerm)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

// 删除目录或者文件

func remove() {
	/* err := os.Remove("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} */

	err := os.RemoveAll("a")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func wd() {
	dir, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir)
	os.Chdir("d:/")
	dir, _ = os.Getwd()
	fmt.Printf("dir: %v\n", dir)

	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

func rename() {
	err := os.Rename("test.txt", "test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func readFile() {
	b, _ := os.ReadFile("test2.txt")
	fmt.Printf("b: %v\n", string(b[:]))
}

func writeFile() {
	os.WriteFile("test2.txt", []byte("hello"), os.ModePerm)
}

func main() {
	// createFile()
	// makeDir()
	// remove()
	// wd()
	// rename()
	// readFile()
	writeFile()
}
 */