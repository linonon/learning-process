package main

import (
	e2 "build-web-with-go/ch02-6-interface/e2"
	"build-web-with-go/ch02-6-interface/e3"
	"build-web-with-go/ch02-6-interface/e4"
	"fmt"
)

func main() {
	e2.E2print()
	fmt.Println("---------------")
	e3.E3print()
	fmt.Println("---------------")
	e4.E4print()
	fmt.Println("---------------")
	e4.E4print2()
}
