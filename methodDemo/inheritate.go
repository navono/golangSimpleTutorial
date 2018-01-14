package methodDemo

import "fmt"

type human struct {
	name  string
	age   int
	phone string
}

type student struct {
	human  //匿名字段
	school string
}

type employee struct {
	human   //匿名字段
	company string
}

//Human定义method
func (h *human) sayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Employee的method重写Human的method
func (e *employee) sayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}
