/*Channels
- are really designed to synchronize data transmission between multiple go routines
- channels is strongly typed
- make(chan int)
*/
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	/*
		ch := make(chan int)
		wg.Add(2)
		//Reciving goroutine
		go func() {
			i := <-ch
			fmt.Println((i))
			wg.Done()
		}()
		//sending routine
		go func() {
			ch <- 42
			wg.Done()
		}()
		wg.Wait()
	*/
	ch := make(chan int, 50) //create buffer of 50 interger
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)
	wg.Wait()
	fmt.Println()
}
