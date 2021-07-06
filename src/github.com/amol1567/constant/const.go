package main

import (
	"fmt"
	//"math"
)

/* Naming convention
   Typed constants
   Untyped constants
   Enumerated constants
   Enumeration expressions
*/
func main() {

	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	//const myConst1 float64 = math.Sin(1.57) //runtime allocation not allow

	//fmt.Printf("%v, %T\n", myConst1, myConst1)

	const a = 42
	var b int16 = 24
	fmt.Printf("%v, %T\n", a+b, a+b)

}
