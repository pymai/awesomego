package istruct

import "fmt"

// 如果结构体的字段类型是 指针、slice 和 map，它们的零值都是 nil，即还没有分配空间
// 如果需要使用这样的字段，需要先 make，才能使用
type Dog struct {
	Name      string
	Age       int
	Scores    [5]float64
	Ptr       *int
	SliceInt  []int
	MapString map[string]string
}

func DogStruct() {
	var dog Dog
	fmt.Println(dog)

	// 会报错 panic: runtime error: index out of range [0] with length 0
	//dog.SliceInt[0] = 100
	//fmt.Println(dog.SliceInt)

	// make 后，值为 [0 0 0 0 0 0 0 0 0 0]
	dog.SliceInt = make([]int, 10)
	// 值为 [100 0 0 0 0 0 0 0 0 0]
	dog.SliceInt[0] = 100
	fmt.Println(dog.SliceInt)

	// 会报错 panic: assignment to entry in nil map
	//dog.MapString["name"] = "foobar"
	//fmt.Println(dog.MapString)

	// make 后，值为 map[]
	dog.MapString = make(map[string]string)
	// 值为 map[name:foobar]
	dog.MapString["name"] = "foobar"
	fmt.Println(dog.MapString)

	// 2021.08.31 为什么这里不会报错？
	num := 10
	dog.Ptr = &num
	fmt.Println(*dog.Ptr)
}
