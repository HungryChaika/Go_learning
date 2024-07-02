package main

import (
	"fmt"
	"io"
)

type phoneReader string
type phoneWriter struct{}
type Writer interface {
	Write(p []byte) (n int, err error)
}

func (ph phoneReader) Read(p []byte) (int, error) {
	count := 0
	for i := 0; i < len(ph); i++ {
		if ph[i] >= '0' && ph[i] <= '9' {
			p[count] = ph[i]
			count++
		}
	}
	return count, io.EOF
}

func (p phoneWriter) Write(bs []byte) (int, error) {
	if len(bs) == 0 {
		return 0, nil
	}
	for i := 0; i < len(bs); i++ {
		if bs[i] >= '0' && bs[i] <= '9' {
			fmt.Println(string(bs[i]))
		}
	}
	fmt.Println()
	return len(bs), nil
}

func main() {
	//test1()
	test2()
}

func test1() {
	phone1 := phoneReader("+1(234)56 7890")
	phone2 := phoneReader("+9-87-65-4321-0")
	buffer1 := make([]byte, len(phone1))
	phone1.Read(buffer1)
	fmt.Println(buffer1)
	fmt.Println(string(buffer1))
	buffer2 := make([]byte, len(phone2))
	phone2.Read(buffer2)
	fmt.Println(buffer2)
	fmt.Println(string(buffer2))
}

func test2() {
	phone1 := []byte("+1(234)56 7890")
	phone2 := []byte("+9-87-65-4321-0")
	writer := phoneWriter{}
	writer.Write(phone1)
	writer.Write(phone2)
}
