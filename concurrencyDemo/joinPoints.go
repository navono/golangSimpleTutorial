package concurrencyDemo

import (
	"fmt"
	"sync"
)

func init() {
	fmt.Println("================= goroutine join point =================")

	// var wg sync.WaitGroup
	// salutation := "hello"
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	salutation = "welcom"
	// }()
	// wg.Wait()
	// // output: welcom
	// fmt.Println(salutation)

	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		// 	// 存在 goroutine 在开始前，循环就结束了，而 `salutation`保存的是最后一个元素的地址。
		//  // 因此会输出3次 `good day`
		// 	fmt.Println(salutation)
		// }()

		// 想要正确输出应该这样：
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
