package main

import (
	"fmt"
)

func test() {
	defer func() {
		err := recover() // 內置函數，可捕獲異常
		if err != nil {
			fmt.Println("err=", err)
			fmt.Println("error sending...")
		}
	}()
	num1 := 100
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}
func main() {
	test()

	fmt.Println("Continue...")
	// defer  panic  recover
}
