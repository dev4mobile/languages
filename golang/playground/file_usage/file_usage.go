package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"playground/util"
)

func WriteFile() {

	d1 := []byte("hello world\n")
	err := ioutil.WriteFile("/tmp/data1", d1, 0644)
	util.Check(err)

	f, err := os.Create("/tmp/data1")
	util.Check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	util.Check(err)
	fmt.Printf("Write %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	util.Check(err)
	fmt.Printf("Write %d bytes\n", n3)
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("Write %d bytes\n", n4)
	w.Flush()
}

func ReadFile() {
	d1, err := ioutil.ReadFile("/tmp/data1")
	util.Check(err)
	fmt.Printf("%s\n", string(d1))

	f, err := os.Open("/tmp/data1")
	util.Check(err)
	defer f.Close()

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	util.Check(err)
	fmt.Printf("Read %d bytes: %s\n", n1, string(b1))

	ret, err := f.Seek(1, 0)
	util.Check(err)
	fmt.Printf("ret: %d\n", ret)
	b2 := make([]byte, 5)
	n2, err := f.Read(b2)
	fmt.Printf("Read %d bytes: %s\n", n2, string(b2))
	f.Seek(0, 0)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	fmt.Printf("Read %d bytes: %s\n", n3, string(b3))

	f.Seek(0, 0)
	r := bufio.NewReader(f)
	b4, err := r.Peek(2)
	fmt.Printf("Read bytes: %s\n", string(b4))
}
