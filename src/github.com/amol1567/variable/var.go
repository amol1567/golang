package main

import (
	"fmt"
	"math"
	"strconv"
)

var b float64 = 12
//var b float32
//b = 12
// b := 12.

func demo() {
	var a = 9
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	demo()
	//fmt.Println(a)

	fmt.Println(b)

	n1 := math.Sqrt(b)
	fmt.Printf("Square root for %.2f is %.2g\n", b, n1)
	fmt.Printf("%v %T", n1, n1)
	fmt.Println()

	n2 := math.Round(n1)
	fmt.Printf("Round is %g\n", n2)
	fmt.Printf("%v %T", n2, n2)
	fmt.Println()

	n3 := math.Ceil(n1)
	fmt.Printf("Ceil is %g\n", n3)
	fmt.Printf("%v %T", n3, n3)
	fmt.Println()

	n4 := math.Floor(n1)
	fmt.Printf("Floor is %g\n", n4)
	fmt.Printf("%v %T\n", n4, n4)

	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	//j = string(i)
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}
