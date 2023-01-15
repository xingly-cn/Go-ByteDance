package main

/**
a 协程发送0-9数字, b 协程计算平方并输出
*/
import "fmt"

func main() {

	src := make(chan int)
	dest := make(chan int, 3)

	// a 协程
	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()

	// b 协程
	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()

	// 主协程
	for i := range dest {
		fmt.Println(i)
	}
}
