package main

import (
	"fmt"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 2; i++ {
		go func(i int) {
			for {
				// a[i]++ // a[i]++ 無法交出控制權
				// runtime.Gosched()
				fmt.Printf("Hello from "+"goroutine %d\n", i)
				time.Sleep(500)
			}
		}(i) // 調用匿名函數, i 是 外部傳入變量
	}
	time.Sleep(time.Minute) //一毫秒後 mian 會結束， mian結束後，go routine就會被kill
	fmt.Println(a)
}
