package e4

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

// 類型斷言
func E4print() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{name: "Dennis", age: 70}

	for i, element := range list {
		if v, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", i, v)
		} else if v, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", i, v)
		} else if v, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", i, v)
		} else {
			fmt.Printf("list[%d] is of a different type\n", i)
		}
	}
}

// 類型斷言 Switch 型
func E4print2() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{name: "Dennis", age: 70}

	for i, element := range list {
		switch v := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", i, v)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", i, v)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", i, v)
		default:
			fmt.Printf("list[%d] is of a different type\n", i)
		}
	}
}
