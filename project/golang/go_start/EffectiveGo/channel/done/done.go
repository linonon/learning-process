package main

import (
	"fmt"
)

func doWorker(i int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("%d %c\n", i, n)
		done <- true
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func creatWorker(i int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(i, w.in, w.done)
	return w
}
func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = creatWorker(i)
	}

	// 第一步，主路徑先跑完 下面的for循環，發10個數據讓 worker打印，發完不管了
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		// <-workers[i].done  // 順序打印
	}
	// 打印完以後就會送 true 到done，沒打印完的時候回阻塞。從而使得這個函數不會執行結束。
	for _, worker := range workers {
		<-worker.done
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		// <-workers[i].done // 順序打印
	}

	for _, worker := range workers {
		<-worker.done
	}
}

func main() {
	chanDemo()
}
