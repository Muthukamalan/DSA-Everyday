package main

func Factorial(n int) int {
	r := 1
	for i := 1; i <= n; i++ {
		r *= i
	}
	return r
}
