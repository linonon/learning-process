package ex1

import "fmt"

func Ex1() {
	c := make(chan int, 2) // 2 -> 1, Error!
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
