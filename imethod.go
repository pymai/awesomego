package main

import "github.com/pymai/awesomego/imethod"

// 方法是作用在指定的数据类型上的（即和指定的数据类型绑定），因此自定义类型也可以有方法，而不仅仅是 struct

func main() {
	imethod.MethodStruct()
}
