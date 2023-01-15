package main

/**
对变量执行2000次+1, 5个协程并发执行 - 线程安全问题
多个协程对临界区访问, 必须加锁

tip：因为不知道子协程多久会完成，所以用sleep是不合适的
*/
import (
	"fmt"
	"sync"
	"time"
)

var (
	x    int64
	lock sync.Mutex // 类比pv操作
)

// 加锁版
func addWithLock() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

// 不加锁
func addWithoutLock() {
	for i := 0; i < 2000; i++ {
		x += 1
	}
}

func main() {
	x = 0
	for i := 0; i < 5; i++ {
		go addWithoutLock()
	}
	time.Sleep(time.Second)
	fmt.Println("未加锁", x)

	x = 0
	for i := 0; i < 5; i++ {
		go addWithLock()
	}
	time.Sleep(time.Second)
	fmt.Println("加锁", x)
}
