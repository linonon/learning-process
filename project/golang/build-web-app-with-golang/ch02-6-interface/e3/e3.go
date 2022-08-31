package e3

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) String() string {
	return "<" + h.name + "-" + strconv.Itoa(h.age) + " years - phone:" + h.phone + ">"
}

// String 重寫
func E3print() {
	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is :", Bob) // 一樣的功能
	fmt.Println("This Human is :", Bob.String())
}
