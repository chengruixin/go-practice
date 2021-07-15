package main

import (
	"fmt"
	"time"
)

func main() {
	
	start := time.Now()
	fibonacci(44)
	end := time.Now()
	
	fmt.Printf("longCalculation took this amount of time: %s\n", end.Sub(start))
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}