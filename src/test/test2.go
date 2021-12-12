package main

import "fmt"

func f() (a int) {

	defer func() {
		a = 100
		fmt.Println("fdfd")
		// defer func ()  {
		// 	a++
		// }()
	}()

	defer func() {
		a = 10
		fmt.Println("fdfd2")
		// defer func ()  {
		// 	a++
		// }()
	}()

	fmt.Println("12")
	return a
}
func main() {
	fmt.Println(f())
}