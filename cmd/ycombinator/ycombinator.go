package ycombinator

// Define a type 'Func' that represents a function taking a parameter of type 'T' and returning a result of type 'U'.
type Func[T, U any] func(T) U

// Define a type 'TagFunc' that represents a function taking a function of type 'Func[T, U]' and returning a function of the same type.
type TagFunc[T, U any] func(Func[T, U]) Func[T, U]

// Define a type 'CombinatorFunc' that represents a function taking a function of type 'CombinatorFunc[T, U]' and returning a function of type 'Func[T, U]'.
type CombinatorFunc[T, U any] func(CombinatorFunc[T, U]) Func[T, U]

// Define a higher-order function 'Y' that takes a 'TagFunc' as an argument and returns a 'Func[T, U]'.
// The 'Y' function essentially implements the Y combinator for functional recursion.
func Y[T, U any](f TagFunc[T, U]) Func[T, U] {
	// Return a function that takes 'self' of type 'CombinatorFunc[T, U]' and returns a 'Func[T, U]'.
	return func(self CombinatorFunc[T, U]) Func[T, U] {
		// Inside the returned function, apply 'f' to a new function that takes 'n' of type 'T' and returns 'U'.
		return f(func(n T) U {
			// Recursively call 'self' with 'n' as the argument to implement recursion.
			return self(self)(n)
		})
	}(func(self CombinatorFunc[T, U]) Func[T, U] {
		// This is the initial self-application, where 'f' is applied to a new function that starts the recursion.
		return f(func(n T) U {
			// Recursively call 'self' with 'n' as the argument to implement recursion.
			return self(self)(n)
		})
	})
}
