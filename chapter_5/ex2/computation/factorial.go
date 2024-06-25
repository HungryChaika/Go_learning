package computation

func Factorial(num int) int {
	if num < 0 {
		return 0
	}
	if num == 0 || num == 1 {
		return 1
	}
	return num * Factorial(num-1)
}
