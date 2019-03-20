// Example provided with help from Jason Waldrip.
// Package work manages a pool of goroutines to perform work.
package work

import "sync"

type Worker interface {
	Task(id int)
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New creates a new work pool
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func(id int) {
			for w := range p.work {
				w.Task(id)
			}
			p.wg.Done()
		}(i)
	}

	return &p
}

// Run submits work to the pool
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown waits for all the goroutines to shutdown
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
