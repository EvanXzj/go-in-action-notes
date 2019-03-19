// This sample program demonstrates how to use an unbuffered
// channel to simulate a game of tennis between two goroutines.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)

	wg.Add(2)

	go player("A", court)
	go player("B", court)

	// Start the set
	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Palyer %s Missed\n", name)
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player.
		court <- ball
	}
}
