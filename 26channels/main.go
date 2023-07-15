package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	myCh := make(chan int)

	wg.Add(2) // for 2 go subroutines

	// receive only
	go func(wg *sync.WaitGroup, myCh <-chan int) {
		val, isChannelOpen := <-myCh // You can choose to print <-myCh directly but this approach helps to differentiate if a 0 value is sent due to the channel being closed or open

		if isChannelOpen {
			fmt.Println(val)
		}
		wg.Done()
	}(wg, myCh)

	// send only channel
	go func(wg *sync.WaitGroup, myCh chan<- int) {
		myCh <- 5
		close(myCh)
		wg.Done()
	}(wg, myCh)

	wg.Wait()
}
