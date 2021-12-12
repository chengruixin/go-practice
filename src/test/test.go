package main

import "fmt"

func main() {
	var start = 1
	var end = 10091
	var sum int
	for i := start; i <= end; i++ {
		sum += i
	}

	fmt.Println((start + end) * (end - start + 1) / 2)
}

