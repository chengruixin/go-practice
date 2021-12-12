package main

import "fmt"

func main() {
	sl := make([]int, 3, 4)

	sl = sl[:len(sl) + 1]
	// sl = sl[:len(sl) + 1]


	fmt.Println(sl, sl[2:2], sl[2:3])
}