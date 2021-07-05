package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pymai/awesomego/ijson"
	"github.com/pymai/awesomego/istruct"
)

// init() 在 main 函数之前执行，可以用来初始化日志格式
func init() {
	log.SetPrefix("[hello.go] ")
	file, err := os.OpenFile("./hello.go.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	// 最后实现的格式
	// [hello.go] 2021/06/01 01:11:04.940030 /Users/khmai/awesomego/hello.go:30: 读取 json 文件
}

func Go() {
	fmt.Println("hello goroutine.")
}

func main() {
	//fmt.Println("Hello World!")

	// 创建 json 文件
	ijson.WriteJson()

	// 读取 json 文件
	//log.Fatalln("读取 json 文件") // 输出信息，并且 "with exit code 1"
	//log.Panicln("读取 json 文件") // 输出信息，并且 "with exit code 2"
	log.Println("读取 json 文件")
	ijson.ReadJson()

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

	ijson.StructToJson()

	ijson.StructToJsonCustom()

	istruct.CompositionStruct()

	foobar := "foobar"
	// ptrVar 是个指针类型变量
	ptrVar := &foobar
	fmt.Println(*ptrVar, *&foobar)
	fmt.Printf("%T\n", ptrVar)

	// 在指针使用上，* 的作用有两个，将 * 放在类型前面表示声明指针类型，将 * 放在变量前面表示解引用操作
	// *string --> 后面跟的是类型，就表示指向 string 的指针类型
	// *ptrVar --> 后面跟的是变量，就表示对指针类型变量进行解引用操作
	var name *string
	name = ptrVar
	fmt.Println(name)
	fmt.Println(*name)
}
