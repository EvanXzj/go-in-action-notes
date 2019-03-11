package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("Recieved: %d\n", i)
	}
	wg.Done()
}

func main() {
	ch := make(chan int)
	go printer(ch)
	wg.Add(1)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()
}
