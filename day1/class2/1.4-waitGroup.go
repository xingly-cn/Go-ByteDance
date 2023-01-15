package main

/**
快速打印
*/
import (
	"fmt"
	"sync"
)

func hello1(i int) {
	fmt.Println("协程", i)
}
func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			defer wg.Done()
			hello1(j)
		}(i)
	}
	wg.Wait() // 与上文的sleep相比, 可以更准确地销毁主线程, 而不是等待一个不确定的时间
}
