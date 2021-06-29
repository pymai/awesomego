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

func StructToJsonCustom() {
	// json 包要求 struct 中的字段必须以大写字母开头，类似 CamelCase 驼峰型命名规范，但有时候需要 snake_case 蛇形命令规范
	// 解决办法：为字段注明标签，使得 json 包在进行编码的时候能够按照标签里的样式修改字段名
	type Location struct {
		// `` 表示原始字符串，json 表示要转换的类型（还有 xml 这些）
		Longitude float64 `json:"longitude_shen"`
		Latitude  float64 `json:"latitue_zhen"`
	}
	shenzhen := Location{114.05, 22.55}
	// json 包中的 Marshal 函数可以将 struct 中的数据转化为 JSON 格式
	bytes, err := json.Marshal(shenzhen)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}
