package main

import "fmt"

func main() {
	capitals := map[string] string {"France":"Paris", "Italy":"Rome", "Japan":"Tokyo" }
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}

	a := make(map[int]string)
	b := make(map[int]string)
	a[1] = "1"
	b[1] = "1"
	r1 := [2]int{0}
	r2 := [2]int{1 : 0}


	fmt.Println(r1 == r2)
	// create map
	mym := make(map[int]string)

	// set map
	mym[2] = "values"
	mym[1] = "values1"
	
	// delete
	delete(mym, 1)
	// traverse
	for k, v := range mym {
		fmt.Println(k, v)
	}

	
}