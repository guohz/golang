/* package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func testCopy() {
	// file1 file2
	r := strings.NewReader("hello world")

	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}

func testCopyBuffer() {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)

	// buf is used here...
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	// ... reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
}

func main() {
	/* r := strings.NewReader("hello world")
	buf := make([]byte, 20)
	r.Read(buf)
	fmt.Printf("string(buf): %v\n", string(buf)) */
	// testCopy()
	testCopyBuffer()
}
 */