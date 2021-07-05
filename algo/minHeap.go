package main

import "fmt"

func main() {
	var arr = []int{}
	arr = push(arr, 4)
	arr = push(arr, 2)
	arr = push(arr, 7)
	arr = push(arr, 9)
	arr = push(arr, 1)
	arr = push(arr, 11)
	arr = push(arr, 99)
	arr = push(arr, 45)
	arr = push(arr, 120)
	fmt.Println(arr)

	fmt.Println(pop(&arr))
	fmt.Println(pop(&arr))
	fmt.Println(pop(&arr))
	fmt.Println(pop(&arr))
	fmt.Println(pop(&arr))
	fmt.Println(pop(&arr))
	fmt.Println(pop(&arr))
	fmt.Println(pop(&arr))
	fmt.Println(pop(&arr))
	
}

func pop(array *[]int) int {
	if len(*array) == 0 {
		panic("arr length is 0")
	}
 
	arr := *array // we should use *array throughout the whole code, but for writing convience, I created a new var which is wasting memory

	fmt.Printf("array addres and copied addr: %p, %p", array, &arr)
	// fmt.Println(&arr)

	if len(arr) == 1 {
		ans := arr[0]
		arr = arr[:0]
		*array = arr
		return ans
	}

	ans := arr[0]
	arr[0] = arr[len(arr) - 1]
	arr = arr[:len(arr) - 1]

	currentIndex := 0

	for true {
		left, right := getLeaves(currentIndex)

		minIdx := getMinIdx(arr, left, right)

		if minIdx == -1 {
			break
		}

		if arr[currentIndex] > arr[minIdx] {
			arr[minIdx], arr[currentIndex] = arr[currentIndex], arr[minIdx]
			currentIndex = minIdx
		} else {
			break
		}

	}

	*array = arr
	return ans
}

func push(arr []int, element int) []int {
	if len(arr) == 0 {
		return append(arr, element)
	}
	
	arr = append(arr, element)
	
	var currentIndex int = len(arr) - 1
	var parentIndex int = getParentIndex(currentIndex)

	for arr[currentIndex] < arr[parentIndex] {
		// swap
		arr[currentIndex], arr[parentIndex] = arr[parentIndex], arr[currentIndex]

		//re assign
		currentIndex = parentIndex

		if currentIndex == 0 {
			break
		} else {
			parentIndex = getParentIndex((currentIndex))
		}

	}

	return arr
}

func getParentIndex(index int) int {
	if index%2 == 0 {
		return index/2 - 1
	}

	return index / 2
}


func getLeaves(parent int) (int, int) {
	return parent * 2 + 1, parent * 2 + 2
}

func getMinIdx(arr []int, left int, right int) int {
	if left >= len(arr) {
		return -1
	}

	if left == len(arr) - 1 {
		return left
	}

	if arr[left] < arr[right] {
		return left
	}

	return right
}