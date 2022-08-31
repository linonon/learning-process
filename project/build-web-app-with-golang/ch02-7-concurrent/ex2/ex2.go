package ex2

import "fmt"

func fibonacci(n int, c chan int) {
	defer close(c)

	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

}

func Ex2() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
