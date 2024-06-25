package main

import "fmt"

type Stream interface {
	read() string
	write(string)
	close()
}

type Qwerty interface {
	qwe()
}

type File struct {
	text string
}

type Folder struct{}

//File methods
func (f *File) read() string {
	return f.text
}

func (f *File) write(text string) {
	f.text = text
	fmt.Println("Запись в файл: ", text)
}

func (f *File) close() {
	fmt.Println("Файл закрыт")
}

//Необходимо наличие всех методов интерфейса для реализации интерфейса (лишние не важны)
func (f *File) qwe() {
	fmt.Println("qwerty")
}

//Folder methods
func (f *Folder) close() {
	fmt.Println("папка закрыта")
}

func main() {
	fmt.Println("Start!")
	myFile := &File{}
	var myFile2 File = File{}
	myFolder := &Folder{}
	writeToStream(myFile, "Hello Go!")
	closeStream(myFile)
	//closeStream(myFolder)		Ошибкат т.к. Folder не реализует Stream
	//writeToStream(myFile2, "Hello Go!")	Ошибка несоответстваия типов File и *File
	myFolder.close()
	myFile2.close()
	//Теперь File реализует два интерфейса
	qwertyPrint(myFile)
}

func writeToStream(stream Stream, text string) {
	stream.write(text)
}

func closeStream(stream Stream) {
	stream.close()
}

func qwertyPrint(qwerty Qwerty) {
	qwerty.qwe()
}
