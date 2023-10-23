# Y Combinator

In the realm of functional programming, the Y combinator serves as a means to formally establish recursive functions within a programming language lacking native support for recursion.

This combinator finds utility in illustrating Curry's paradox, a concept that exposes the inherent inconsistency of untyped lambda calculus as a deductive system. The Y combinator achieves this by enabling an anonymous expression to represent zero or even multiple values, thus challenging mathematical logic.

When applied to a single-variable function, the Y combinator typically leads to non-termination. However, more intriguing outcomes emerge when it is applied to functions with two or more variables. These additional variables can serve as counters or indices, resulting in a function that mimics the behavior of a "while" or "for" loop in an imperative language.

In this usage, the Y combinator facilitates basic recursion. In the lambda calculus, direct self-reference within a function's body using its name is impossible. Instead, recursion is achieved by obtaining the same function as an argument and employing that argument for the recursive call, rather than relying on the function's name, as is the case in languages with built-in recursion support. The Y combinator serves as a compelling exemplar of this programming approach.

# Go Implementation

The code defines a higher-order function 'Y' that implements the Y combinator for functional recursion. It takes a 'TagFunc' and uses it to create a recursive function of type 'Func[T, U]'. This code is a powerful tool in functional programming for creating recursive functions.


Define a type 'Func' that represents a function taking a parameter of type 'T' and returning a result of type 'U'.
```
type Func[T, U any]           func(T) U
```

Define a type 'TagFunc' that represents a function taking a function of type 'Func[T, U]' and returning a function of the same type.
```
type TagFunc[T, U any]        func(Func[T, U]) Func[T, U]
```

Define a type 'CombinatorFunc' that represents a function taking a function of type 'CombinatorFunc[T, U]' and returning a function of type 'Func[T, U]'
```
type CombinatorFunc[T, U any] func(CombinatorFunc[T, U]) Func[T, U]
```

Define a higher-order function 'Y' that takes a 'TagFunc' as an argument and returns a 'Func[T, U]'.
The 'Y' function essentially implements the Y combinator for functional recursion.

```
func Y[T, U any](f TagFunc[T, U]) Func[T, U] {
	return func(self CombinatorFunc[T, U]) Func[T, U] {
		return f(func(n T) U {
			return self(self)(n)
		})
	}(func(self CombinatorFunc[T, U]) Func[T, U] {
		return f(func(n T) U {
			return self(self)(n)
		})
	})
}

```

### Full Code

```
type Func[T, U any]           func(T) U
type TagFunc[T, U any]        func(Func[T, U]) Func[T, U]
type CombinatorFunc[T, U any] func(CombinatorFunc[T, U]) Func[T, U]

func Y[T, U any](f TagFunc[T, U]) Func[T, U] {
	return func(self CombinatorFunc[T, U]) Func[T, U] {
		return f(func(n T) U {
			return self(self)(n)
		})
	}(func(self CombinatorFunc[T, U]) Func[T, U] {
		return f(func(n T) U {
			return self(self)(n)
		})
	})
}

```


# Examples

Bellow we have some classical recursive problems implemented using the Y combinator

### Factorial
```
var factorial_tag = func(recurse Func[int, int]) Func[int, int] {
	return func(n int) int {
		if n == 0 {
			return 1
		}
		return n * recurse(n-1)
	}
}

// usage

func main() {
    fac := Y(factorial_tag)
    n := 5
    val := fac(n)

    fmt.Printf("Factorial of %d: %d\n", n, val)
}
```

### Fibonacci Sequence
```
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
    
	fib := Y(fib_tag)
	n = 6
	val = fib(n)
	fmt.Printf("Fibonacci of %d: %d\n", n, val)
}

```