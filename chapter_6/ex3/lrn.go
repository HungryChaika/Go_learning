package main

import "fmt"

type Writer interface {
	write(string)
}

type Reader interface {
	read() string
}

type ReaderWriter interface {
	Reader
	Writer
}

type File struct {
	text string
}

func (f *File) read() string {
	return f.text
}

func (f *File) write(text string) {
	f.text = text
}

func main() {
	myFile := &File{}
	readToStream(myFile)
	writeToStream(myFile, "Congratulation")
	readToStream(myFile)
}

func writeToStream(writer ReaderWriter, text string) {
	writer.write(text)
}

func readToStream(reader ReaderWriter) {
	fmt.Println("text: ", reader.read())
}
