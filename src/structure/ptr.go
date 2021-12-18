package structure

import "fmt"

type Gender int

const (
	male Gender = iota
	female
)

type Person struct {
	name   string
	age    int
	gender Gender
}

func fortest() {
	var aman *Person = &Person{"ruixin", 2, male}
	var awoman *Person = new(Person)
	awoman.name = "weiwei"
	awoman.age = 1
	awoman.gender = 2

	var another Person = Person{name: "unknown", gender: 23}
	fmt.Println(aman, awoman)

	awoman = &another

	fmt.Println(another, awoman)

	another = Person{name: "known"}
	fmt.Println(aman, awoman, another)

	aman = &another

	fmt.Println(aman, awoman, another)

	fmt.Println(*aman == *awoman, *aman == another)

	var mm = new(map[int]int)
	(*mm)[1] = 2
	fmt.Println(mm)
}
