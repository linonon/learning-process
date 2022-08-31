package e2

import "fmt"

type Human struct {
	Name  string
	Age   int
	Phone string
}

type Student struct {
	Human
	School string
	Loan   float64
}

type Employee struct {
	Human
	Company string
	Money   float64
}

//Human对象实现Sayhi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.Name, h.Phone)
}

// Human对象实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

// Employee重载Human的Sayhi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.Name,
		e.Company, e.Phone) //此句可以分成多行
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

//
func E2print() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	var i Men

	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	i = tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)

	x[0], x[1], x[2] = paul, sam, mike

	for _, v := range x {
		v.SayHi()
	}

}
