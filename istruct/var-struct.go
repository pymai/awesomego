package istruct

import "fmt"

type Db struct {
	Name    string
	Version float64
}

func VarStruct() {
	// 创建结构体实例的 4 种方式

	// 方式 1
	var d1 Db
	d1.Name = "mysql"
	d1.Version = 5.6
	fmt.Println(d1)

	// 方式 2
	d2 := Db{}
	d2.Name = "postgre"
	d2.Version = 11.2
	fmt.Println(d2)

	// 方式 3
	// var d3 *Db = new(Db)
	var d3 = new(Db)
	// (*d3).Name = "redis"
	d3.Name = "redis"
	// (*d3).Version = 3.2
	d3.Version = 3.2
	fmt.Println(*d3)

	// 方式 4
	// var d4 *Db = &Db{}
	var d4 = &Db{}
	// (*d4).Name = "oracle"
	d4.Name = "oracle"
	// (*d4).Version = 11.2
	d4.Version = 11.2
	fmt.Println(*d4)

}
