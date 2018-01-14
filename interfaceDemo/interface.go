package interfaceDemo

import (
	"fmt"
	"reflect"
	"strconv"
)

// 所有的类型都实现了空interface
// 空interface在我们需要存储任意类型的数值的时候相当有用，
// 因为它可以存储任意类型的数值。它有点类似于C语言的void*类型

type human struct {
	name  string
	age   int
	phone string
}

type student struct {
	human  //匿名字段
	school string
	loan   float32
}

type employee struct {
	human   //匿名字段
	company string
	money   float32
}

//human实现SayHi方法
func (h human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//human实现Sing方法
func (h human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//employee重载human的SayHi方法
func (e employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Interface men被human,student和employee实现
// 因为这三个类型都实现了这两个方法
type men interface {
	SayHi()
	Sing(lyrics string)
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h human) String() string {
	return "<" + h.name + " - " + strconv.Itoa(h.age) + "  years -  ✆ " + h.phone + ">"
}

type Element interface{}
type List []Element

func init() {
	fmt.Println("================= interfaceDemo =================")

	mike := student{human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := student{human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := employee{human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := employee{human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义men类型的变量i
	var i men

	//i能存储student
	i = mike
	fmt.Println("This is Mike, a student:")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储employee
	i = tom
	fmt.Println("This is tom, an employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice men
	fmt.Println("Let's use a slice of men and see what happens")
	x := make([]men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}

	fmt.Println("This Human is: ", x[0])

	// 判断 interface{} 的类型
	list := make(List, 3)
	list[0] = x[0]
	list[1] = 1
	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(student); ok {
			fmt.Printf("list[%d] is a Student and its value is %s\n", index, value)
		}

		// switch value := element.(type) {
		// case int:
		// 	fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		// case student:
		// 	fmt.Printf("list[%d] is a Student and its value is %s\n", index, value)
		// }
	}

	v := reflect.ValueOf(list[0])
	fmt.Println("type: ", v.Type())
}
