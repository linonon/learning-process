package e

import "fmt"

type Human struct {
	Name  string
	Age   int
	Phone string
}

// 繼承
type Student struct {
	Human
	School string
}

type Employee struct {
	Human
	Company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I'm %s you can call me on %s.\n", h.Name, h.Phone)
}

// 重寫 Human 的 SayHi
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I'm %s, I work at %s. Call me on %s\n", e.Name, e.Company, e.Phone)
}

func E3print() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}
