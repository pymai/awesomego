package istruct

import "fmt"

type WebSite struct {
	Name    string
	Address string
}

func PointerStruct() {
	w1 := WebSite{"qq", "www.qq.com"}
	fmt.Println(w1)

	w2 := &w1
	fmt.Println(*w2)

	// 这里和我之前理解的不一样，之前以为 w2 没有自己的地址
	// 输出的结果：w2 有自己的地址，w2 的值为 w1 的地址，也就是 w2 指向 w1
	// w1 的地址 0x1400007a080
	// w2 的地址 0x1400000e030  w2 的值0x1400007a080
	fmt.Printf("w1 的地址 %p\n", &w1)
	fmt.Printf("w2 的地址 %p\t w2 的值%p\n", &w2, w2)

	// . 的优先级高，先执行 w2.name 获得 string qq ，然后执行 * string qq，所以会报错
	// Invalid indirect of 'w2.Name' (type 'string')
	//fmt.Println(*w2.Name)
}
