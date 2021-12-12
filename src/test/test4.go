package main
import "fmt"

func fp(a *[3]int) { 
	fmt.Println(a) 
}

func main() {
	for i := 0; i < 3; i++ {
		fp(&[3]int{i, i * i, i * i * i})
	}

	ar := [2]int{1, 2}

	var sl []int = ar[:]

	sl[0] = -1

	fmt.Println(ar, sl)


	sl2 := []int{1, 2}


	sl2 = appendSlice(sl2)
	sl2 = appendSlice(sl2)

	fmt.Println(sl2)

	editSlice(sl2)
	editSlice(sl2)
	fmt.Println("edited", sl2)

	s3 := append(sl2, 222)

	fmt.Println("apped", sl2, s3)
}


func appendSlice(sl []int) []int {
	return append(sl, 2)
}

func editSlice(sl []int) {
	sl[0] = sl[0] -1
}
func arr(a [2]int) {
	a[0] = -1
	fmt.Println(a)
}