package ichannel

import (
	"fmt"
	"time"
)

func SleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore ...")

	// 发送操作会等待直到另一个 goroutine 尝试对该通道进行接收操作为止
	// 		执行发送操作的 goroutine 在等待期间将无法执行其他操作
	// 		未在等待通道操作的 goroutine 仍然可以继续自由的运行

	// 向通道发送值： c <- id
	c <- id
}
