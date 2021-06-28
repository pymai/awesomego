package ijson

import (
	"encoding/json"
	"fmt"
	"os"
)

func StructToJson() {
	type Location struct {
		// 结构体成员首字母要大写，不然会输出 {}
		Longitude, Latitude float64
	}

	shenzhen := Location{114.05, 22.55}
	// json 包中的 Marshal 函数可以将 struct 中的数据转化为 JSON 格式
	bytes, err := json.Marshal(shenzhen)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 如果输出小写字母的 JSON 数据？
	fmt.Println(string(bytes))
}
