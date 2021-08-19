package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	tty, err := os.OpenFile("./test.txt", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	var r io.Reader
	r = tty

	var w io.Writer

	w = r.(io.Writer) // Reader与Writer的pair中的type相同。

	w.Write([]byte("HELLO THIS IS A TEST!!!\n"))
}
