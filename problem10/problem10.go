package problem10

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b, res := 0, 1, 0
	for i := 2; i <= n; i++ {
		res = a + b
		a, b = b, res
	}
	return res
}
