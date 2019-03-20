// This sample program demonstrates how to use the work package
// to use a pool of goroutines to get work done.
package main

import (
	"github.com/EvanXzj/go-in-action-notes/chapter07/patterns/work"
	"log"
	"sync"
	"time"
)

// names provides a se of names to display
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task(id int) {
	log.Println("GID", id, "print:", m.name)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(5 * len(names))

	for i := 0; i < 5; i++ {
		for _, name := range names {
			np := namePrinter{
				name,
			}

			go func() {
				p.Run(&np)

				wg.Done()
			}()
		}
	}

	wg.Wait()

	p.Shutdown()
}
