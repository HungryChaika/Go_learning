package main

import (
	"fmt"
	"reflect"
)

type mile uint
type kilometer uint
type abbreviation_operation func(int, int) int
type psevdonim_int = int
type struct_elems struct {
	method func(int, int) int
	num    int
}
type nested_struct struct {
	number int
}
type main_struct struct {
	elem nested_struct
}
type link_struct struct {
	number int
	next   *link_struct
}

func main() {
	fmt.Println("Start")
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//test7()
	//test8()
	test9()
}

func test1() {
	var elem mile = 100
	fmt.Println("elem before: ", elem)
	elem += 50
	fmt.Println("elem after: ", elem)
	fmt.Println("////////////////////////////////////")
}

func test2() {
	var (
		distance1 mile      = 6
		distance2 kilometer = 7
	)
	distance(distance1)
	fmt.Println(reflect.TypeOf(distance2))
	distance(mile(distance2)) //Тип отличный от mile не примет
}

func test3() {
	action(4, 5, add)
	action_short(5, 6, add)
	fmt.Println("////////////////////////////////////")
}

func test4() {
	var elem psevdonim_int
	fmt.Println(reflect.TypeOf(elem))
}

func test5() {
	var elem struct_elems = struct_elems{add, 45}
	elem0 := struct_elems{}
	fmt.Println("elem: ", elem)
	fmt.Println("elem0: ", elem0)
	fmt.Println("add res:")
	fmt.Println(elem.method(1, 3))
	fmt.Println("num before:", elem.num)
	var elemPointer *struct_elems = &elem
	(*elemPointer).num = 33
	fmt.Println("num after:", elem.num)
}

func test6() {
	var (
		elem1      *struct_elems = &struct_elems{method: add, num: 12}
		elem2      *struct_elems = new(struct_elems)
		numPointer *int          = &elem1.num
	)
	fmt.Println(elem2)
	fmt.Println("num before:", elem1.num)
	*numPointer = 37
	fmt.Println("num after:", elem1.num)
}

func test7() {
	var nested_elem nested_struct = nested_struct{4}
	var main_elem main_struct = main_struct{nested_elem}
	fmt.Println(main_elem)
	fmt.Println(main_elem.elem.number)
}

func test8() {
	var first link_struct = link_struct{number: 1}
	second := link_struct{number: 2}
	var third link_struct = link_struct{number: 3}
	first.next = &second
	second.next = &third
	var current *link_struct = &first
	for current != nil {
		fmt.Println(current.number)
		current = current.next
	}
}

func test9() {

}

/*
===================================================================================================================================================
===================================================================================================================================================
*/

func distance(distance mile) {
	fmt.Println("Расстояние: ", distance, " миль")
}

func action(x, y int, operation func(int, int) int) {
	result := operation(x, y)
	fmt.Println(result)
}

func action_short(x, y int, operation abbreviation_operation) {
	result := operation(x, y)
	fmt.Println(result)
}

func add(x, y int) int {
	return x + y
}
