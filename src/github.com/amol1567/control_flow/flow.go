/*Control Flow
-if, else if, else
-switch
-for statements

*/
package main

import (
	"fmt"
	"math"
)

func main() {
	if true {
		fmt.Println("this is true.")
	}

	number := 50
	guess := 50
	fmt.Println(number <= guess, number >= guess, number != guess)

	myNum := 0.123
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else if myNum > 0.001 {
		fmt.Println("These are different")
	} else {
		fmt.Println("Okay")
	}

	num := 5
	switch num {
	case 1, 5, 9:
		fmt.Println("one five or nine")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other number")
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough //pass
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than 20")
	}

	var j interface{} = 1
	switch j.(type) {
	case int:
		fmt.Println("j is an int")
		//break
		fmt.Println("this will print too")
	case float64:
		fmt.Println("j is a float64")
	case string:
		fmt.Println("j is string")
	default:
		fmt.Println("j is another type")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 5; i < 5; i, j = i+1, j-1 {
		fmt.Println(i, j)
	}

	t := 0
	for t < 5 {
		fmt.Println(t)
		t++
	}

	//s := []int{1, 2, 3}
	s := "hello world"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}

}
