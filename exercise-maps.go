package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	var i int
	

	for _,j := range words {
		i = m[j]
		i++
		m[j] = i
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
