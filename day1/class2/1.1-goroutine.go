package main

/**
使用协程快速（并行）输出
*/
import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("协程", i)
}

func main() {

	for i := 0; i < 5; i++ {
		go func(j int) {
			hello(j)
		}(i)
	}
	time.Sleep(time.Second)
}
