/* Arrays
- Collection of items with same type
-Fixed size
-Declaration styles
	- a := [3]int{1,2,3}
	- a := [...]int{1,2,3}
	- var a [3]int
Slicing:
-A slice has both a length and a capacity.
-The length of a slice is the number of elements it contains.
-The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
-via make function
	- a := make([]int, 10) // create slice with capacity and length == 10
	- a := make([]int, 10, 100) // create slice with length == 10 and capacity == 100
*/
package main

import "fmt"

func main() {

	//grades := [3]int{92, 24, 67}
	grades := [...]int{92, 24, 67}
	fmt.Printf("Grades: %v\n", grades)

	var students [5]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "amol"
	students[2] = "akshat"
	students[1] = "amit"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Students #1: %v\n", students[1])
	fmt.Printf("No of students: %v\n", len(students))

	var identityMatrix = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(identityMatrix)

	a := [...]int{1, 2, 3}
	b := a
	// b := &a // b is pointing to same data instead of creating new
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	n1 := []int{1, 2, 3}
	fmt.Printf("Length: %v\n", len(n1))
	fmt.Printf("Capacity: %v\n", cap(n1))

	n2 := make([]int, 3) //built-in
	fmt.Println(n2)
	fmt.Printf("Length: %v\n", len(n2))
	fmt.Printf("Capacity: %v\n", cap(n2))

	n3 := []int{}
	fmt.Println(n3)
	fmt.Printf("Length: %v\n", len(n3))
	fmt.Printf("Capacity: %v\n", cap(n3))
	n3 = append(n3, 1, 2, 3, 4, 5)
	n3 = append(n3, []int{2, 3, 4}...) //spliting the array
	fmt.Println(n3)
	fmt.Printf("Length: %v\n", len(n3))
	fmt.Printf("Capacity: %v\n", cap(n3))

	//stack
	n4 := []int{1, 2, 3, 4, 5}
	n5 := n4[1:]
	n6 := n4[:len(n4)-1]
	n7 := append(n4[:2], n4[3:]...)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
}
