package main

import (
	"github.com/pymai/awesomego/igoroutine"
)

func main() {
	// 会打印 ...snore...
	//go igoroutine.SleepyGopher() // 分支线路，运行了 3 秒
	//time.Sleep(4 * time.Second)  // 主线路，运行了 4 秒
	// main() 函数走完后，无论 goroutine 是否走完，程序都会停止

	// 启动 goroutine，只需在调用前面加一个 go 关键字
	// 不会打印 ...snore...
	go igoroutine.SleepyGopher() // 分支线路，要运行 3 秒
	// 但是 main() 函数会直接执行完（可能不到 1 秒），所以没有等到 goroutine 的结果，程序就退出了

}
