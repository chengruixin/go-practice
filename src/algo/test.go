package main

import "fmt"

//[8,6,9]
// [1,7,5]
// 3
func main() {
	arr1 := []int{6, 7}
	arr2 := []int{6, 0, 4}
	// fmt.Println(getMaxSubsequence(arr1, 1))
	fmt.Println(maxNumber(arr1, arr2, 5))
	// fmt.Println(getMaxSubsequence(arr, 2))
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
    len1, len2 := len(nums1), len(nums2)
	ans := make([]int, k)

	start := 0
	if k > len2 {
		start = k - len2
	}
	// fmt.Println(start)
	for i := start ; i <= k && i <= len1; i++ {
		fmt.Println(i)

		sub1 := getMaxSubsequence(nums1, i)
		sub2 := getMaxSubsequence(nums2, k - i)
		curMax := merge(sub1, sub2)
		fmt.Println(curMax, i, k - i)
		if leftLarger(curMax, 0, ans, 0) {
			ans = curMax
		}
	}

	
	return ans
}	

func getMaxSubsequence(arr []int, max int) []int {
	var stack []int

	for i, item := range arr {
		for len(stack) > 0 &&
			item > stack[len(stack)-1] &&
			len(arr)-i-1 >= max {
			stack = stack[:len(stack)-1]
			max++
		}

		if max > 0 {
			stack = append(stack, item)
			max--
		}

	}

	return stack
}

func merge(arr1, arr2 []int) []int {
	var ans []int

	ptr1, ptr2 := 0, 0

	for ptr1 < len(arr1) || ptr2 < len(arr2) {
		if leftLarger(arr1, ptr1, arr2, ptr2) {
			ans = append(ans, arr1[ptr1])
			ptr1++
		} else {
			ans = append(ans, arr2[ptr2])
			ptr2++
		}

		if ptr1 >= len(arr1) {
			for ptr2 < len(arr2) {
				ans = append(ans, arr2[ptr2])
				ptr2++
			}
		}

		if ptr2 >= len(arr2) {
			for ptr1 < len(arr1) {
				ans = append(ans, arr1[ptr1])
				ptr1++
			}
		}
	}
	
	return ans
}

func leftLarger(arr1 []int, ptr1 int, arr2 []int, ptr2 int) bool {
	if ptr1 >= len(arr1) || ptr2 >= len(arr2) {
		if ptr2 >= len(arr2) {
			return true
		}
		return false
	}
	n := math.MinIn
	if arr1[ptr1] == arr2[ptr2] {
		return leftLarger(arr1, ptr1+1, arr2, ptr2+1)
	} else if arr1[ptr1] > arr2[ptr2] {
		return true
	} else {
		return false
	}
}