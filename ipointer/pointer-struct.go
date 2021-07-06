package ipointer

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) birthday() {
	p.age++
}

func PointerStruct() {
	foobar := person{name: "foobar", age: 18}
	// 在变量通过点标记法进行调用的时候，自动使用 & 取得变量的内存地址
	foobar.birthday()
	fmt.Printf("%+v\n", foobar)
}

//func birthday(p *person) {
//	p.age++
//}
//
//func PointerStruct() {
//	foobar := person{name: "foobar", age: 18}
//	birthday(&foobar)
//	fmt.Printf("%+v\n", foobar)
//}

// map 在被赋值或者被作为参数传递的时候不会被复制，map 就是一种隐式指针
// slice 在指向数据元素的时候也使用了指针
// 当 slice 被直接传递至函数或方法时， slice 的内部指针就可以对底层数据进行修改
// interface 也可以直接调用指针方法
