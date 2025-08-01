package main

import (
	"fmt"
)

func recursion(n int) int {
	if n == 0 {
		return 1
	}

	return n * recursion(n-1)
}

func main() {
	fmt.Println(recursion(3))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			fmt.Print(n)
			return n
		}
		fmt.Println(n)
		return fib(n-1) + fib(n-2)
	}

	fmt.Print("\n")
	fmt.Println(fib(7))

}
