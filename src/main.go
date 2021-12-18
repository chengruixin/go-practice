package main

import (
	"fmt"
	math "ruixin/mymath"
	"ruixin/structure"
)

func main() {
	math.Print(math.Add(1, 2))
	m := structure.NewMatric(1, 2)
	fmt.Println(m)
	var a *structure.TagType = new(structure.TagType)
	a.Shit = 111

	fmt.Println(a)

}
