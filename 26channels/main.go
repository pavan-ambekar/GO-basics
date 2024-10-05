package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to channels in golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)
	wg.Add(2)
	// Receive only
	go func(myCh <-chan int, wg *sync.WaitGroup) {
		// close(myCh)
		val, isChanelOpen := <-myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)

	// send only
	go func(myCh chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 1
		wg.Done()
	}(myCh, wg)

	defer wg.Wait()
}
