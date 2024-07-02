package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type phoneReader string

func (p phoneReader) Read(bs []byte) (int, error) {
	count := 0
	for i := 0; i < len(p); i++ {
		if p[i] >= '0' && p[i] <= '9' {
			bs[count] = p[i]
			count++
		}
	}
	return count, io.EOF
}

type Person struct {
	name   string
	age    int
	weight float64
}

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//test7()
	//test8()
	//test9()
	//test10()
	//test11()
	//test12()
	//test13()
	test14()
}

func test1() {
	file, err := os.Create("myText.txt") //os.Open("NAME")
	//f1, err := os.OpenFile("Text1.txt", os.O_RDONLY, 0666)
	//f2, err := os.OpenFile("Text2.txt", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("ERROR!!! ", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Println(file.Name())
}

func test2() {
	var (
		text      string = "Simple text, what must destroy the Galaxy!"
		byte_text []byte = []byte("Be calm, this text is weaker than the previous...")
	)

	file, err := os.Create("Surprise.txt")
	if err != nil {
		fmt.Println("ERROR!!! ", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(text)
	file.Write(byte_text)
	fmt.Println("Done!")
}

func test3() {
	file, err := os.Open("myText.txt")
	if err != nil {
		fmt.Println("ERROR!!! ", err)
		os.Exit(1)
	}
	defer file.Close()
	data := make([]byte, 8)
	for {
		n, err := file.Read(data)
		if err == io.EOF { //Если конец файла
			break
		}
		time.Sleep(2 * time.Second)
		fmt.Println("\nn: ", n)
		fmt.Println(string(data[:n]))
	}
}

func test4() {
	file, err := os.Open("myText.txt")
	if err != nil {
		fmt.Println("ERROR!!! ", err)
		os.Exit(1)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}

func test5() {
	phone1 := phoneReader("+1(777)560 89")
	io.Copy(os.Stdout, phone1)
}

func test6() {
	file, err := os.Create("FmtTest.txt")
	if err != nil {
		fmt.Println("ERROR! ", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fprint(file, "Today ")
	fmt.Fprintln(file, "is good weather")
}

func test7() {
	human := Person{
		name:   "Jack",
		age:    12,
		weight: 100.987,
	}
	file, err := os.Create("PersonBook.dat")
	if err != nil {
		fmt.Println("ERROR! ", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fprintf(file,
		"%-10s %-10d %-10.3f\n",
		human.name, human.age, human.weight)
}

func test8() {
	human := Person{
		name:   "Yolo",
		age:    85,
		weight: 7865.112,
	}
	fmt.Printf("%-10s %-10d %-10.3f\n",
		human.name, human.age, human.weight)
	fmt.Printf("Hello  ")
	fmt.Println("cold!")
}

func test9() {
	var filename string = "test9.txt"
	writeData(filename)
	readData(filename)
}

func test10() {
	var filename = "test10.dat"
	writeForTest10(filename)
	readForTest10(filename)
}

func test11() {
	var filename = "test11.dat"
	var (
		people = []Person{
			{"Petya", 18, 90.12},
			{"Vasya", 23, 45.345},
			{"Katya", 41, 71.6789},
		}
		name   string
		age    int
		weight float64
	)
	crFile, crErr := os.Create(filename)
	if crErr != nil {
		fmt.Println("CREATE ERROR BEACH!\n", crErr)
		os.Exit(1)
	}
	for _, p := range people {
		fmt.Fprintf(crFile, "%s %d %.2f\n", p.name, p.age, p.weight)
	}
	crFile.Close()

	//=============================

	opFile, opErr := os.Open(filename)
	if opErr != nil {
		fmt.Println("OPEN ERROR BEACH!\n", opErr)
		os.Exit(1)
	}
	for {
		_, opErr = fmt.Fscanf(opFile, "%s %d %f\n", &name, &age, &weight)
		if opErr != nil {
			if opErr == io.EOF {
				break
			} else {
				fmt.Println("OPEN ERROR BEACH!\n", opErr)
				os.Exit(1)
			}
		}
		fmt.Printf("%-8s %-8d %-8.2f\n", name, age, weight)
	}
	opFile.Close()
}

func test12() {
	var (
		name string
		age  int
	)
	fmt.Print("Введите имя: ")
	fmt.Fscan(os.Stdin, &name)
	fmt.Print("Введите возраст: ")
	fmt.Fscan(os.Stdin, &age)
	fmt.Println(name, age)
	/*
		========================================
	*/
	fmt.Print("Введите имя и возраст: ")
	fmt.Scan(&name, &age)
	fmt.Println(name, age)
}

func test13() {
	var rows []string = []string{
		"Hello Go!",
		"Welcome to Golang",
	}
	file, err := os.Create("test13.dat")
	if err != nil {
		fmt.Println("ERROR!", err)
		os.Exit(1)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, row := range rows {
		writer.WriteString(row)
		writer.WriteString("\n")
	}
	writer.Flush()
}

func test14() {
	file, err := os.Open("test13.dat")
	if err != nil {
		fmt.Println("ERROR!", err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("ERROR!", err)
				return
			}
		}
		fmt.Print(line)
	}
}

/*
====================================================================================================================
====================================================================================================================
*/

func writeData(filename string) {
	var (
		name string = "Vasya"
		age  int    = 777
	)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("CREATE ERROR BEACH!\n", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fprintln(file, name)
	fmt.Fprintln(file, age)
}

func readData(filename string) {
	var (
		name string
		age  int
	)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("OPEN ERROR BEACH!\n", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fscanln(file, &name)
	fmt.Fscanln(file, &age)
	fmt.Println(name, age)
}

func writeForTest10(filename string) {
	var person Person = Person{name: "Ivan", age: 90, weight: 1.23}
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("CREATE ERROR BEACH!\n", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fprintf(file, "%s %d %.2f\n", person.name, person.age, person.weight)
}

func readForTest10(filename string) {
	var (
		name   string
		age    int
		weight float64
	)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("OPEN ERROR BEACH!\n", err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = fmt.Fscanf(file, "%s %d %f\n", &name, &age, &weight)
	if err != nil {
		fmt.Println("READ ERROR BEACH!\n", err)
		os.Exit(1)
	}
	fmt.Printf("%-8s %-8d %-8.2f\n", name, age, weight)
}
