/*Concurrency
- goroutines abstract the concept of thread
- ability of the application to work on multiple things at the same time.
- goroutines is use to create concurrent and parallel execution paths in our applications.
- when using aninymous functions, pass data as local variables.
Synchronization - how to interact with one another
	- use sync.WaitGroup to wait for groups of goroutines to complete
	- use sync.Mutex and sync.RWMutex to protect data access
Parallelism -
	- by default, Go will use CPU threads equal to available cores
	- Change with runtime.GOMAXPROCS
-Don't create goroutines in libaries: Let consumer control concurrency

*/
package main

import (
	"fmt"
	"runtime"
	"sync"
	//"time"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {

	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	/*
		go sayHello()
		time.Sleep(100 * time.Millisecond)
		//fmt.Print()
	*/

	/*
		var msg = "Hello"
		go func() {
			fmt.Println(msg)
		}()
		msg = "Good Bye"
		time.Sleep(100 * time.Millisecond)
	*/

	var msg = "HEllo"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "GoodBye"
	wg.Wait()

	for i := 0; i < 5; i++ {
		runtime.GOMAXPROCS(100)
		wg.Add(2)
		m.RLock()
		go sayHello1()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Println("Hello")
}

func sayHello1() {
	//m.RLock()
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}
