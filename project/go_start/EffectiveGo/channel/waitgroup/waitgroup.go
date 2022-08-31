package main

import (
	"fmt"
	"sync"
)

func doWorker(i int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("%d %c\n", i, n)
		wg.Done()
	}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func creatWorker(i int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(i, w.in, wg)
	return w
}
func chanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i] = creatWorker(i, &wg)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo()
}
