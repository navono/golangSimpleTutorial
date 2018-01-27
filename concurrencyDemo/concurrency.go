package concurrencyDemo

import (
	"fmt"
	"time"
)

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func fibonacciV1(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func testClose() {
	fc := make(chan int, 10)
	go fibonacciV1(cap(fc), fc)
	for i := range fc {
		fmt.Println(i)
	}
}

func fibonacciV2(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		case <-time.After(2 * time.Second):
			println("timeout")
			return
			// default就是当监听的channel都没有准备好的时候，默认执行的（select不再阻塞等待channel）
		}
	}
}

func testSelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		// quit <- 0
	}()
	fibonacciV2(c, quit)
}

// 默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，
// 而不需要显式的lock。所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。其次，任何发送（ch<-5）将会被阻塞，直到数据被读出
func init() {
	fmt.Println("================= concurrencyDemo =================")

	a := []int{7, 2, 8, -9, 4, 0}

	// 无缓冲channel
	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	// ch := make(chan type, value)
	// 当 value = 0 时，channel 是无缓冲阻塞读写的，
	// 当value > 0 时，channel 有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入。

	bc := make(chan int, 3) //修改2为1就报错，修改2为3可以正常运行
	bc <- 1
	bc <- 2
	// fmt.Println(<-bc)
	// fmt.Println(<-bc)

	testSelect()
}
