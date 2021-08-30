package istruct

import "fmt"

// 定义一个 Cat 结构体，将 Cat 的各个字段/属性信息，放入到 Cat 结构体进行管理
// 结构体是个值类型，不是引用类型
type Cat struct {
	Name  string
	Age   int
	Color string
}

func CatStruct() {
	// 创建一个 Cat 的变量
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	fmt.Println("cat1 =", cat1)

}
