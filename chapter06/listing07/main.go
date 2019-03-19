package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Allocate 2 logical processor for the scheduler to use.
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	//
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTermination Program")
}
