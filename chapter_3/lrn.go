package main

import "fmt"

func main() {
	fmt.Println("Start")
	//test1()
	//test2()
	test3()
}

func test1() {
	var elem int = 8
	var link_elem *int
	link_elem = &elem
	fmt.Println("Address: ", link_elem)
	fmt.Println("Value : ", *link_elem)
	fmt.Println("elem before: ", elem)
	*link_elem = 14
	fmt.Println("elem after: ", elem)
	if link_elem != nil {
		fmt.Println("link_elem is not null, it's", *link_elem)
	}
	fmt.Println("////////////////////////////////////")
}

func test2() {
	elem := new(int)
	fmt.Println("value before: ", *elem)
	*elem = 4
	fmt.Println("value after: ", *elem)
	elemSquaring(elem)
	fmt.Println("value after x2: ", *elem)
	var elemNotLink int = 5
	fmt.Println("elemNotLink before: ", elemNotLink)
	elemSquaring(&elemNotLink)
	fmt.Println("elemNotLink after: ", elemNotLink)
	fmt.Println("////////////////////////////////////")
}

func test3() {
	var elem *int = returnLingElem(80)
	fmt.Println("Address: ", elem)
	fmt.Println("Value : ", *elem)
	fmt.Println("////////////////////////////////////")
}

/*
===================================================================================================================================================
===================================================================================================================================================
*/

func elemSquaring(elem *int) {
	*elem = (*elem) * (*elem)
}

func returnLingElem(elem int) *int {
	link_elem := new(int)
	*link_elem = elem
	return link_elem
}
