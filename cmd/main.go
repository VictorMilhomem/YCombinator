package main

import (
	"fmt"

	. "github.com/VictorMilhomem/YCombinator/cmd/ycombinator"
)

var factorial_tag = func(recurse Func[int, int]) Func[int, int] {
	return func(n int) int {
		if n == 0 {
			return 1
		}
		return n * recurse(n-1)
	}
}

var fib_tag = func(recurse Func[int, int]) Func[int, int] {
	return func(n int) int {
		if n <= 0 {
			return 0
		} else if n == 1 {
			return 1
		}
		return recurse(n-1) + recurse(n-2)
	}
}

func main() {
	// factorial example
	fac := Y(factorial_tag)
	n := 5
	val := fac(n)

	fmt.Printf("Factorial of %d: %d\n", n, val)

	// fibonacci example

	fib := Y(fib_tag)
	n = 6
	val = fib(n)
	fmt.Printf("Fibonacci of %d: %d\n", n, val)
}
