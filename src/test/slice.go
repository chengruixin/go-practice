package main

import "fmt"

func main() {
	b:= []byte{'g', 'o', 'l', 'a', 'n', 'g'}

	fmt.Println(string(b[1:4]), string(b[:2]), string(b[2:]), string(b[:]))


	slice := make([]int, 10)
	fmt.Println(slice)

	s2 := new([1]int)
	var s3 [1]int

	s3 = *s2

	fmt.Println(s2, s3)

	s := make([]byte, 5)

	fmt.Println(cap(s), len(s))

	s = s[2:4]

	fmt.Println(cap(s), len(s))


	sll1 := []rune{'a', 'b', 'c', '说'}

	sll2 := sll1[2:]

	sll2[1] = '吧'

	fmt.Println(string(sll1), string(sll2))


	sll3 := []int{1, 2, 3}

	for i := range sll3 {
		fmt.Println(i)
	}
}

func mor() (int, int, int) {
	return 1, 2, 3
}