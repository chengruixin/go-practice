package main

import "fmt"

func main() {
	var slice = []int{1, 2}

	slice = changeSlice(slice)

	fmt.Println("outer", slice)

}

func digest(a ...int) func(added int) {
	fmt.Println(a, "is digested")
	sum := 0
	for _, v := range a {
		sum += v
	}
	return func(added int) {
		sum += added
		fmt.Println(sum)
	}
}

func changeSlice(slice []int) []int {
	// s := *slice
	return append(slice, slice[0])
	

}