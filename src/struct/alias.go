package main

import "fmt"

type evals struct {
    a float32
	b string
}

type nr evals   // alias type

func main() {
    a := evals{a: 5.0}
    b := nr{a: 5.0}
    // var i float32 = b   // compile-error: cannot use b (type nr) as type float32 in assignment
    // var i = float32(b)  // compile-error: cannot convert b (type nr) to type float32
    // var c number = b    // compile-error: cannot use b (type nr) as type number in assignment
    // needs a conversion:
	c := b
	c.b = "sdf"
    var d = c
    fmt.Println(a, b, d)

	d = nr(a)
	a.b = "sdsffsdf"	
    fmt.Println(a, b, d)

}