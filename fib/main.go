package main

import "fmt"

func main() {
	fmt.Println("fib is: ", fib(1, make(map[int]int)))
	fmt.Println("fib is: ", fib(5, make(map[int]int)))
	fmt.Println("fib is: ", fib(10, make(map[int]int)))
	fmt.Println("fib is: ", fib(50, make(map[int]int)))
}

func fib(n int, memo map[int]int) int {
	if n < 3 {
		return 1
	}
	value, ok := memo[n]
	if ok {
		return value
	}

	memo[n] = fib(n-1, memo) + fib(n-2, memo)

	return memo[n]
}
