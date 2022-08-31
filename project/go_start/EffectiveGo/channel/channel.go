package main

import (
	"fmt"
	"time"
)

func worker(i int, c chan int) {

	// for {
	// 	n, ok := <-c
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("%d %c\n", i, n)
	// }

	// 同樣效果

	for n := range c {
		fmt.Printf("%d %c\n", i, n)
	}
}

// chan<- 只能發數據到chan
// <-chan 只能從chan收數據
func creatWorker(i int) chan<- int {
	c := make(chan int)
	go worker(i, c)
	return c
}
func chanDemo() {
	//  ver c chan int // c == nil 沒法用
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		// channels[i] = make(chan int)
		// go worker(i, channels[i])

		// 創建10個worker
		channels[i] = creatWorker(i)
	}

	// c <- 1 // 發了沒人收會 deadlock
	// c <- 2 // 發數據到c

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 4) // 緩衝區大小是3
	go worker(0, c)
	c <- '1'
	c <- '2'
	c <- '3'
	c <- '4'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int) // 緩衝區大小是3
	go worker(0, c)
	c <- '1'
	c <- '2'
	c <- '3'
	c <- '4'
	close(c)
	time.Sleep(time.Millisecond)
}
func main() {
	chanDemo()
	// bufferedChannel()
	// channelClose()

}
