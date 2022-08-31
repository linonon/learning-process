package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	g := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) * time.Millisecond)
			g <- i
			i++
		}
	}()
	return g
}
func worker(i int, c chan int) {
	for n := range c {
		time.Sleep(1 * time.Second)
		fmt.Printf("Recived %d , n: %d.\n", i, n)
	}
}

func creatWorker(i int) chan<- int {
	c := make(chan int)
	go worker(i, c)
	return c
}
func main() {
	c1, c2 := generator(), generator()
	w := creatWorker(0)

	var values []int
	tmten := time.After(10 * time.Second)
	tmtick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("Time Out")
		case <-tmtick:
			fmt.Println("Queue", len(values))
		case <-tmten:
			fmt.Println("bye")
			return

		}
	}

}
