package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int = 0

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
	test13()
}

func test1() {
	fmt.Println("START")
	for i := 0; i < 10; i++ {
		go factorial(i)
	}
	fmt.Scanln() //Ожидание ввода пользователя, что-бы функция factorial, выполняемая параллельно main успела выполниться 10 раз (цикл)
	fmt.Println("END")
}

func test2() {
	fmt.Println("START")
	for i := 1; i < 10; i++ {
		go func(n int) {
			result := 1
			for j := 1; j <= n; j++ {
				result *= j
			}
			fmt.Println("result ", n, ": ", result)
		}(i)
	}
	fmt.Scanln()
	fmt.Println("END")
}

func test3() {
	chint := make(chan int)
	fmt.Println("START")
	go func() {
		fmt.Println("Go routine STARTS")
		chint <- 5
	}()
	fmt.Println(<-chint)

	fmt.Println("END")
}

func test4() {
	chint := make(chan int)
	fmt.Println("START")
	fmt.Println("Num is 4")
	go square(chint)
	chint <- 4
	fmt.Println(<-chint)
	fmt.Println("END")
}

func test5() {
	fmt.Println("START")
	channel := make(chan int, 4)
	channel <- 1
	channel <- 2
	channel <- 3
	channel <- 4
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println("END")
}

func test6() {
	fmt.Println("START")
	channel := make(chan int, 4)
	channel <- 890
	fmt.Println("len: ", len(channel))
	fmt.Println("cap: ", cap(channel))
	fmt.Println("END")
}

func test7() {
	fmt.Println("START")
	inch := make(chan int, 4)
	sum(inch, 4, 5)
	fmt.Println(<-inch)
	fmt.Println("END")
}

func test8() {
	fmt.Println("START")
	fmt.Println(<-createChannel(5))
	fmt.Println("END")
}

func test9() {
	fmt.Println("START")
	chint := make(chan int, 3)
	chint <- 6
	chint <- 8
	chint <- 9
	fmt.Println(<-chint)
	close(chint)
	//chint <- 999 #Ошибка, канал закрыт
	for i := 0; i < cap(chint); i++ {
		if val, oppened := <-chint; oppened {
			fmt.Println(val)
		} else {
			fmt.Println("Channel is closed")
		}
	}
	fmt.Println("END")
}

func test10() {
	results := make(map[int]int)
	structCh := make(chan struct{})
	fmt.Println("START")
	go multipeFactorial(structCh, results, 6)
	<-structCh
	for i, v := range results {
		fmt.Println(i, " it's ", v)
	}
	fmt.Println("END")
}

func test11() {
	fmt.Println("START")
	channel1 := make(chan int)
	channel2 := make(chan int)
	go factorialForDataflow(channel1, 6)
	for {
		num, oppened := <-channel1
		if !oppened {
			break
		}
		fmt.Println(num)
	}
	fmt.Println("PRE END")
	go factorialForDataflow(channel2, 4)
	for num := range channel2 {
		fmt.Println(num)
	}
	fmt.Println("END")
}

func test12() {
	fmt.Println("START")
	channel := make(chan bool)
	var mutex sync.Mutex
	for i := 1; i < 5; i++ {
		go work(channel, i, &mutex)
	}
	for i := 1; i < 5; i++ {
		<-channel
	}
	fmt.Println("END")
}

func test13() {
	fmt.Println("START")
	var wg sync.WaitGroup
	wg.Add(2)
	myWork := func(id int) {
		defer wg.Done()
		fmt.Println("Goroutine started ", id)
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine finished ", id)
	}
	go myWork(1)
	go myWork(2)
	wg.Wait()
	fmt.Println("END")
}

/*
======================================================================================================================
======================================================================================================================
*/

func factorial(num int) {
	if num < 0 {
		fmt.Println("Invalid value!")
		return
	}
	var result int = 1
	if num == 0 {
		fmt.Println("result ", num, ": ", result)
		return
	}
	for i := 1; i <= num; i++ {
		result *= i
	}
	fmt.Println("result ", num, ": ", result)
}

func square(ch chan int) {
	num := <-ch
	fmt.Println("square 4 is")
	ch <- num * num
}

func sum(ch chan<- int, x, y int) {
	ch <- x + y
}

func createChannel(num int) chan int {
	channel := make(chan int)
	go func() {
		channel <- num
	}()
	return channel
}

func multipeFactorial(ch chan struct{}, res map[int]int, num int) {
	defer close(ch)
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
		res[i] = result
	}
}

func factorialForDataflow(ch chan int, num int) {
	defer close(ch)
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
		ch <- result
	}
}

func work(ch chan bool, num int, mutex *sync.Mutex) {
	mutex.Lock()
	counter = 0
	for i := 1; i < 5; i++ {
		counter++
		fmt.Println("Go routine: ", num, " - ", counter)
	}
	mutex.Unlock()
	ch <- true
}
