/*
Method
	- a function that takes a receiver
*/

package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		sayMessage("Hello Go!", i)
	}
	fmt.Println()

	sum("The sum is ", 1, 2, 3, 4, 5)

	s := sumret(1, 2, 3, 4, 5, 6)
	fmt.Println("The sum is ", *s)

	//Anonymous
	func() {
		fmt.Println("Hello Go!")
	}()

	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	//method
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()

	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

func sum(msg string, values ...int) {

	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

func sumret(values ...int) *int {

	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

//method
type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
