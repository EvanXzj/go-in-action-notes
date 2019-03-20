// This sample program demonstrates how to use a channel to
// monitor the amount of time the program is running and terminate
// the program if it runs too long
package main

import (
	"github.com/EvanXzj/go-in-action-notes/chapter07/patterns/runner"
	"log"
	"os"
	"time"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")

	r := runner.New(timeout)

	// add the tasks to be run
	r.Add(createTask(), createTask(), createTask())

	// Run the tasks an handle the result
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Termination due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
