package imethod

import "fmt"

type Coffee struct {
	name string
}

// 这里的小 c 是传值，和函数调用一样，除非 Coffee 使用指针类型
// 输出结果：
// getName 函数里面的 c.name =  latte
// MethodStruct 函数里面的 c.name =  cappuccino
func (c Coffee) getName() {
	c.name = "latte"
	fmt.Println("getName 函数里面的 c.name = ", c.name)
}

// 输出结果：
// getName 函数里面的 c.name =  latte
// MethodStruct 函数里面的 c.name =  latte
//func (c *Coffee) getName() {
//	c.name = "latte"
//	fmt.Println("getName 函数里面的 c.name = ", c.name)
//}

func MethodStruct() {
	var c Coffee
	c.name = "cappuccino"
	c.getName()
	fmt.Println("MethodStruct 函数里面的 c.name = ", c.name)
}
