/* package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func test1() {
	// r := strings.NewReader("hello world")
	f, _ := os.Open("a.txt")
	defer f.Close()
	r2 := bufio.NewReader(f)
	s, _ := r2.ReadString('\n')
	fmt.Printf("s: %v\n", s)
}

func test2() {
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	br := bufio.NewReader(s)
	p := make([]byte, 10)

	for {
		n, err := br.Read(p)
		if err == io.EOF {
			break
		} else {
			fmt.Printf("string(p): %v\n", string(p[0:n]))
		}
	}
}

func test3() {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)

	c, _ := br.ReadByte()
	fmt.Printf("%c\n", c)

	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)

	br.UnreadByte() //
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
}

func test4() {
	s := strings.NewReader("ABC\nDEF\r\nGHI\r\nGHI")
	br := bufio.NewReader(s)

	w, isPrefix, _ := br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)
}

func test5() {
	s := strings.NewReader("ABC,DEF,GHI,JKL")
	br := bufio.NewReader(s)

	w, _ := br.ReadSlice(',')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadSlice(',')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadSlice(',')
	fmt.Printf("%q\n", w)
}

func test6() {
	s := strings.NewReader("ABC DEF GHI JKL")
	br := bufio.NewReader(s)

	w, _ := br.ReadBytes(' ')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadBytes(' ')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadBytes(' ')
	fmt.Printf("%q\n", w)
}

func test7() {
	s := strings.NewReader("ABC DEF GHI JKL")
	br := bufio.NewReader(s)

	w, _ := br.ReadString(' ')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadString(' ')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadString(' ')
	fmt.Printf("%q\n", w)
}

func test8() {
	s := strings.NewReader("hello golang")
	br := bufio.NewReader(s)
	// b := bytes.NewBuffer(make([]byte, 0))

	// File Reader Writer

	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0777)
	defer f.Close()
	br.WriteTo(f)
	// fmt.Printf("%s\n", b)
}

func test9() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0777) // File Reader Writer

	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString("hello world!")
	w.Flush()
}

func test10() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.WriteString("123456789")
	c := bytes.NewBuffer(make([]byte, 0))
	bw.Reset(c)
	bw.WriteString("456")
	bw.Flush()
	fmt.Println(b)
	fmt.Println(c)
}

func test11() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	fmt.Println(bw.Available()) // 4096
	fmt.Println(bw.Buffered())  // 0

	bw.WriteString("ABCDEFGHIJKLMN")
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)

	bw.Flush()
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)
}

func test12() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	s := strings.NewReader("123")
	br := bufio.NewReader(s)
	rw := bufio.NewReadWriter(br, bw)
	p, _ := rw.ReadString('\n')
	fmt.Println(string(p)) //123
	rw.WriteString("asdf")
	rw.Flush()
	fmt.Println(b)
}

func test13() {
	s := strings.NewReader("ABC DEF GHI JKL")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords) // 以空格作为分隔符进行分割
	for bs.Scan() {
		fmt.Println(bs.Text())
	}
}

func test14() {
	s := strings.NewReader("Hello 世界！")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanRunes)
	for bs.Scan() {
		fmt.Printf("%s ", bs.Text())
	}
}

func main() {
	// test2()
	// test3()

	// test4()
	// test5()

	// test6()
	// test7()

	// test8()

	// test9()
	// test10()
	// test11()
	// test12()
	// test13()
	test14()

}
*/