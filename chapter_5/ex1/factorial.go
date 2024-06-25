package main

func factorial(num int) int {
	if num < 0 {
		return 0
	}
	if num == 0 || num == 1 {
		return 1
	}
	return num * factorial(num-1)
}
