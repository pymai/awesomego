package main

import (
	"github.com/pymai/awesomego/itemplate"
)

func main() {
	//itemplate.DotTemplate()
	itemplate.RangeTemplate()
	//server := http.Server{Addr: "0.0.0.0:8888"}
	// 2021.07.28 这行什么意思？ --> 第一个参数是 uri，第二个参数是处理这个 uri 的函数
	//http.HandleFunc("/", itemplate.TemplateIndex)
	//server.ListenAndServe()
}
