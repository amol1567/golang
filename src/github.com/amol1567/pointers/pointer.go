/*Pointer
 *int --> pointer to integer
 *b --> dereferencing (at runtime fetch the value from memory)

 */

package main

import "fmt"

func main() {

	var a int = 42
	var b *int = &a
	fmt.Println(&a, b)

	i, j := 42, 2701

	p := &i // point to i
	fmt.Println(p)
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	a1 := [3]int{1, 2, 3}
	b1 := &a1[0]
	c1 := &a1[1]
	fmt.Printf("%v %p %p\n", a1, b1, c1)

	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)
	(*ms).foo = 42 //ms.foo = 42
	// ms = &myStruct{foo: 42}
	fmt.Println(ms)
	fmt.Println((*ms).foo) //fmt.Println(ms.foo)

}

type myStruct struct {
	foo int
}
