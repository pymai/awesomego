package main

import (
	"fmt"

	"github.com/pymai/awesomego/ichannel"
)

/*
创建通道用 make 函数，并指定其传输数据的类型，比如 c := make(chan int)

使用左箭头操作符 <- 向通道发送值 或 从通道接收值
	向通道发送值： c <- 99
	从通道接收值： r := <- c

发送操作会等待直到另一个 goroutine 尝试对该通道进行接收操作为止
	执行发送操作的 goroutine 在等待期间将无法执行其他操作
	未在等待通道操作的 goroutine 仍然可以继续自由的运行

执行接收操作的 goroutine 将等待直到另一个 goroutine 尝试向该通道进行发送操作为止
*/

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go ichannel.SleepyGopher(i, c)
	}
	// 如果 for i := 0; i < 4; i++ ，那么会有一个通道的值没有被接收，程序就退出了
	/*
		...  3  snore ...
		...  2  snore ...
		...  4  snore ...
		...  0  snore ...
		...  1  snore ...
		gopher  3  has finished sleeping
		gopher  2  has finished sleeping
		gopher  4  has finished sleeping
		gopher  0  has finished sleeping
	*/

	for i := 0; i < 5; i++ {
		// 2021.07.17 这部分有用到 goroutine 吗？
		// 执行接收操作的 goroutine 将等待直到另一个 goroutine 尝试向该通道进行发送操作为止
		// 从通道接收值： gopherID := <-c
		gopherID := <-c
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}
