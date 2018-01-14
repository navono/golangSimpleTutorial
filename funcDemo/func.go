package funcDemo

import "fmt"

// 声明了一个函数类型
type testInt func(int) bool

func isEven(a int) bool {
	return a%2 == 0
}

func isOdd(a int) bool {
	return !isEven(a)
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func init() {
	fmt.Println("================= concurrencyDemo =================")

	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("Odd: ", filter(slice, isOdd))
	fmt.Println("Even: ", filter(slice, isEven))
}
