package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(maxNum)

	fmt.Println("请输入一个数字")
	reader := bufio.NewReader(os.Stdin) // 输入流, cin >>
	for {
		input, err := reader.ReadString('\n') // 从输入流读一行
		if err != nil {
			fmt.Println(err)
			return
			continue
		}
		input = strings.TrimSuffix(input, "\r\n") // 去掉换行符

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("你猜的数字是：", guess)
		switch {
		case guess == num:
			fmt.Println("你猜对啦")
			return
		case guess < num:
			fmt.Println("你猜小了")
		case guess > num:
			fmt.Println("你猜大了")
		}
	}
}
