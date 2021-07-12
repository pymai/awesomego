package main

import (
	"fmt"

	"github.com/pymai/awesomego/ipointer"
)

func main() {
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

	ipointer.PointerStruct()
}
