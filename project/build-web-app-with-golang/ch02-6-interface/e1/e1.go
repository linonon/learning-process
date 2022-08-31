package e

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
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.Name, h.Phone)
}

// Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

//Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee重载Human的Sayhi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.Name,
		e.Company, e.Phone) //此句可以分成多行
}

//Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float64) {
	s.Loan += amount // (again and again and...)
}

//Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float64) {
	e.Money -= amount // More vodka please!!! Get me through the day!
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float64)
}

type ElderGent interface {
	SayHi()
	Sing(lyrics string)
	SpendSalary(amount float64)
}
