package main

import (
	"log"
	"os"

	"github.com/pymai/awesomego/ijson"
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

func main() {
	// 创建 json 文件
	ijson.WriteJson()

	// 读取 json 文件
	//log.Fatalln("读取 json 文件") // 输出信息，并且 "with exit code 1"
	//log.Panicln("读取 json 文件") // 输出信息，并且 "with exit code 2"
	log.Println("读取 json 文件")

	ijson.ReadJson()

	ijson.StructToJson()

	ijson.StructToJsonCustom()
}
