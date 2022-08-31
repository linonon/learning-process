package main

import (
	"fmt"
	"testing"
)

// 編寫測試單元
func TestAddUpper2(t *testing.T) {
	// 調用
	// res := addUpper(10)
	// if res != 55 {
	// 	t.Fatalf("result is wrong(%v), correct answer is :%v.", res, 55)
	// }

	// t.Logf("result is correct")
	tests := []struct {
		n   int
		ans int
	}{
		{10, 55},
		{5, 15},
		{20, 210},
		{30, 455}, // 錯的
		{80, 3240},
		{100, 5050},
		{10000, 50005000},
	}

	for _, tt := range tests {
		res := addUpper2(tt.n)
		if res != tt.ans {
			// t.Fatalf("correct ans should be %v, now is %v", res, tt.ans)
			fmt.Printf("Wrong, correct ans should be %v, now is %v\n", res, tt.ans)
		} else {
			fmt.Println(tt.n, " is correct.")
		}
	}
}

func TestGetSub(t *testing.T) {
	// 調用
	res := sub(10, 3)
	if res != 7 {
		t.Fatalf("result is wrong(%v), correct answer is :%v.", res, 55)
	}

	t.Logf("sub()'s result is correct")
}

func BenchmarkAddUpper2(b *testing.B) {
	n := 10000
	ans := 50005000

	for i := 0; i < b.N; i++ {
		actual := addUpper2(n)
		if actual != ans {
			b.Errorf("error")
		}
	}
}
