package main

import (
	"fmt"
)

func Go() {
	fmt.Println("hello goroutine.")
}

// 2021.07.03 为什么其他文件也可以有 main() 函数
func main() {
	//fmt.Println("Hello World!")

	// 直接运行，看不到 Go() 函数的输出
	// 因为 Go() 函数还没运行完，main() 函数可能就退出了
	//go Go()

	// time.Sleep 运行，可以看到 Go() 函数的输出
	// 这是投机取巧的做法，Go() 不一定在 2 秒运行完
	// 所以需要一种机制来确认，当 Go() 运行完了，才去通知 main()，让 main() 退出，这就需要用到 channel 了
	//go Go()
	//time.Sleep(2 * time.Second)

	c := make(chan bool)
	// 这是匿名函数
	go func() {
		fmt.Println("hello goroutine.")
		c <- true
	}()
	// <-c: channel 等待 c 的值，所以 main() 阻塞了
	// 等 go 匿名函数运行到 c <- true，channel 拿到了 true 的值，才不阻塞，main() 就退出了
	<-c

}
