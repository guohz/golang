/* package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func testTrans() {
	var i int = 100
	var b byte = 10
	b = byte(i)
	i = int(b)

	var s string = "hello world"
	b1 := []byte{1, 2, 3}
	s = string(b1)
	b1 = []byte(s)
}

func testContains() {
	s := "duoke360.com"
	b := []byte(s)
	b1 := []byte("duoke360")
	b2 := []byte("DuoKe360")

	strings.Contains("hello world", "hello")

	b3 := bytes.Contains(b, b1)
	fmt.Printf("b3: %v\n", b3)
	b3 = bytes.Contains(b, b2)
	fmt.Printf("b3: %v\n", b3)
}

func testCount() {
	//Count
	s := []byte("hellooooooooo")
	sep1 := []byte("h")
	sep2 := []byte("l")
	sep3 := []byte("o")
	fmt.Println(bytes.Count(s, sep1)) //1
	fmt.Println(bytes.Count(s, sep2)) //2
	fmt.Println(bytes.Count(s, sep3)) //9
}

func testRepeat() {
	//Repeat
	b := []byte("hi")
	fmt.Println(string(bytes.Repeat(b, 1))) //hi
	fmt.Println(string(bytes.Repeat(b, 3))) //hihihi
}

func testReplace() {
	//Replace
	s := []byte("hello,world")
	old := []byte("o")
	news := []byte("ee")
	fmt.Println(string(bytes.Replace(s, old, news, 0)))  //hello,world
	fmt.Println(string(bytes.Replace(s, old, news, 1)))  //hellee,world
	fmt.Println(string(bytes.Replace(s, old, news, 2)))  //hellee,weerld
	fmt.Println(string(bytes.Replace(s, old, news, -1))) //hellee,weerld
}

func testRunes() {
	//Runes
	s := []byte("你好世界")
	r := bytes.Runes(s)
	fmt.Println("转换前字符串的长度: ", len(s)) //12
	fmt.Println("转换后字符串的长度: ", len(r)) //4
}

func testJoin() {
	//Join
	s2 := [][]byte{[]byte("你好"), []byte("世界")}
	sep4 := []byte(",")
	fmt.Println(string(bytes.Join(s2, sep4))) //你好,世界
	sep5 := []byte("#")
	fmt.Println(string(bytes.Join(s2, sep5))) //你好#世界
}

func testBuffer() {
	var b bytes.Buffer
	fmt.Printf("b: %v\n", b)
	var b1 = bytes.NewBufferString("hello")
	fmt.Printf("b1: %v\n", b1)
	var b2 = bytes.NewBuffer([]byte("hello"))
	fmt.Printf("b2: %v\n", b2)
}

func testBuffer2() {
	var b bytes.Buffer
	n, _ := b.WriteString("hello")
	fmt.Printf("n: %v\n", n)
	fmt.Printf("b: %v\n", string(b.Bytes()))
}

func testBuffer3() {
	var b = bytes.NewBufferString("hello world")
	b1 := make([]byte, 2)
	for {
		n, err := b.Read(b1)
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("b1: %v\n", string(b1[0:n]))
	}

}

func testReader() {

	data := "123456789"
	//通过[]byte创建Reader
	re := bytes.NewReader([]byte(data))
	//返回未读取部分的长度
	fmt.Println("re len : ", re.Len())
	//返回底层数据总长度
	fmt.Println("re size : ", re.Size())
	fmt.Println("------------")

	buf := make([]byte, 2)
	for {
		//读取数据
		n, err := re.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}

	fmt.Println("------------")

	//设置偏移量，因为上面的操作已经修改了读取位置等信息
	re.Seek(0, 0)
	for {
		//一个字节一个字节的读
		b, err := re.ReadByte()
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}
	fmt.Println("------------")

	re.Seek(0, 0)
	off := int64(0)
	for {
		//指定偏移量读取
		n, err := re.ReadAt(buf, off)
		if err != nil {
			break
		}
		off += int64(n)
		fmt.Println(off, string(buf[:n]))
	}
}

func main() {
	// testCount()
	// testRepeat()
	// testReplace()
	// testRunes()
	// testJoin()
	// testBuffer()
	// testBuffer2()
	// testBuffer3()
	testReader()
}
 */