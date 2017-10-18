package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	currFib, nextFib, nextNextFib := 0, 0, 1
	
	return func() int {
		currFib = nextFib
		nextFib = nextNextFib
		nextNextFib = currFib + nextFib
		return currFib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}