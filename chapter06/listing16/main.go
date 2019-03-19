// 这个示例程序展示如何使用互斥锁来
// 定义一段需要同步访问的代码临界区
// 实现资源的同步访问
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter

			// 这句还有用吗？goroutine A 切换到 goroutine B， B 碰到mutex.Lock()又切换回A?
			runtime.Gosched()

			value++
			counter = value
		}
		mutex.Unlock()
	}
}
