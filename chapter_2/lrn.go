package main

import "fmt"

var (
	hello                      string  = "h\tel\nl\"o \\wor\rld"
	first, second                      = "ttr", 45
	num_with_floating_point_32 float32 = 7.123456788888888888888888888888888
	num_with_floating_point_64 float32 = 7.123456788888888888888888888888888
)

const PI float64 = 3.1415926535897932384626433832795028841971693993751058209749445923078164
const (
	a = 1
	b
	c
	d = 2
	e
	f = 3
	g = 4
)

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
	//test14()
	//test15()
	test16()
}

func test1() {
	sym := 'G'
	fmt.Println(sym)
	fmt.Println(num_with_floating_point_32)
	fmt.Println(num_with_floating_point_64)
	fmt.Println("////////////////////////////////////")
}

func test2() {
	fmt.Println(first, second, "all")
	fmt.Println(PI)
	fmt.Println(a, b, c, d, e, f, g)
	fmt.Println("////////////////////////////////////")
}

func test3() {
	var (
		t1 int = 8 & 1
		t2 int = 9 & 1
	)
	fmt.Println(t1, t2)
	if t1 == 1 {
		fmt.Println("Нечётный")
	} else {
		fmt.Println("Чётный")
	}
	fmt.Println("////////////////////////////////////")
}

func test4() {
	var nums1 [4]int = [4]int{1, 2, 3}
	nums2 := [4]int{1, 2, 3}
	nums3 := [...]int{1, 2, 3}
	fmt.Println(nums1, nums2, nums3)
	fmt.Println("////////////////////////////////////")
}

func test5() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("////////////////////////////////////")
}

func test6() {
	var (
		arr1 [3]string = [3]string{"qwerty", "asdfgh", "zxcvbn"}
		arr2 [3]string = [3]string{"123", "456", "789"}
		arr3 [3]int    = [3]int{1, 2, 3}
	)
	for index, value := range arr1 {
		fmt.Println(index, value)
	}
	for _, value := range arr2 {
		fmt.Println(value)
	}
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	fmt.Println("////////////////////////////////////")
}

func test7() {
	fmt.Println(mega_add(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(mega_add(1, 2, 3, 4, 5))
	fmt.Println(simple_add(5, 6))
	fmt.Println("////////////////////////////////////")
}

func test8() {
	NAME, AGE := interesting_add(1, 50, "Petya", "Ivanov")
	fmt.Println(NAME+": ", AGE)
	fmt.Println("////////////////////////////////////")
}

func test9() {
	var elem1 func(int, int) (int, int) = simple_add
	elem2 := mega_add
	fmt.Println(elem1(5, 2))
	elem1 = multiply
	fmt.Println(elem1(5, 2))
	fmt.Println(elem2(4, 5, 6))
	fmt.Println("////////////////////////////////////")
}

func test10() {
	fmt.Println(big_action(4, 3, multiply))
	fmt.Println("////////////////////////////////////")
}

func test11() {
	myfunc1 := selectfunc("add")
	fmt.Println(myfunc1(3, 3))
	myfunc2 := selectfunc("mult")
	fmt.Println(myfunc2(3, 3))
	myfunc3 := selectfunc("")
	fmt.Println(myfunc3(2, 2))
	fmt.Println("////////////////////////////////////")
}

func test12() {
	fmt.Println(small_action(5, 5, func(x, y int) int { return x + y }))
	fmt.Println("////////////////////////////////////")
}

func test13() {
	defer fmt.Println("Is this the end, really?")
	defer fmt.Println("We are nearing the end.")
	for i := 10; i >= 0; i-- {
		fmt.Println(i)
	}
	fmt.Println("////////////////////////////////////")
}

func test14() {
	if true {
		panic("Oh noo, it's TRUUUE!!!")
	}
	fmt.Println("unattainable height")
	fmt.Println("////////////////////////////////////")
}

func test15() {
	var (
		srez_or_slize1       = []int{0, 1, 2, 3}
		srez_or_slize2 []int = make([]int, 4)
	)
	fmt.Println("srez_or_slize1:", srez_or_slize1)
	fmt.Println("cap srez_or_slize1:", cap(srez_or_slize1))
	fmt.Println("srez_or_slize2:", srez_or_slize2)
	srez_or_slize1 = append(srez_or_slize1, 4)
	srez_or_slize1 = append(srez_or_slize1, 5)
	fmt.Println("srez_or_slize1:", srez_or_slize1)
	fmt.Println("len srez_or_slize1:", len(srez_or_slize1))
	fmt.Println("cap srez_or_slize1:", cap(srez_or_slize1))
	fmt.Println("sub 2- srez_or_slize1:", srez_or_slize1[2:])
	fmt.Println("sub 2-len(srez_or_slize1) srez_or_slize1:", srez_or_slize1[2:6]) //6 is len(srez_or_slize1)
	fmt.Println("sub 2-8 srez_or_slize1:", srez_or_slize1[2:8])
	srez_or_slize1 = append(srez_or_slize1[:3], srez_or_slize1[5:]...)
	fmt.Println("del 3 srez_or_slize1:", srez_or_slize1)
	fmt.Println("////////////////////////////////////")
}

func test16() {

}

/*
===================================================================================================================================================
===================================================================================================================================================
*/

func mega_add(nums ...int) int {
	var sum int = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func simple_add(x, y int) (z, w int) {
	z = x + y
	w = x - y
	return
}

func interesting_add(x, y int, name, lastname string) (string, int) {
	sum, fullname := x+y, lastname+name
	return fullname, sum
}

func multiply(x, y int) (z, w int) {
	z = x * y
	w = x / y
	return
}

func big_action(x, y int, operation func(int, int) (int, int)) (int, int) {
	return operation(x, y)
}
func small_action(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func selectfunc(func_num string) func(int, int) (int, int) {
	switch func_num {
	case "add":
		return simple_add
	case "mult":
		return multiply
	default:
		return func(x, y int) (int, int) {
			return x, y
		}
	}
}
