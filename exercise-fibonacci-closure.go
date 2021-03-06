package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib, nextFib, nextNextFib := 0, 0, 1
	
	return func() int {
		fib = nextFib
		nextFib = nextNextFib
		nextNextFib = fib + nextFib
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
