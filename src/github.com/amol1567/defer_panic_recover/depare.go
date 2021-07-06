/*Defer -- works as LIFO
	-Used to delay execution of a statement until function exits
	-Useful to group "open" and "close" functions together
		-Be careful in loops
	-Run in LIFO(Last In, First Out) order
	-Arguments evaluated at the time defer is executed, not at time of called functon execution
Panic -- exception
	-panic execute after defer statement excute
	-occur when program can not continue at all
	-Function will stop executing
		-Deferred functions will still fire
	-If nothing handles panic, program will exit
Recover -- use inside deferred function
	-Used to recover from panic
	-Current function will not attempt to coninue, bu higher functions in call stack will
*/

package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("middle") //execute after function finish its execution
	fmt.Println("end")

	/*a, b := 1, 0
	ans := a / b
	fmt.Println(ans)
	*/

	fmt.Println("start")
	defer fmt.Println(("this was defered"))
	panic("something bad happened")
	//fmt.Println("end")

	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			panic(err)
		}
	}()
	panic("something bad happened")

}
