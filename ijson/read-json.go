package ijson

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJson() {
	path, err := os.Getwd()
	if err != nil {
		// 如果 return 语句使用在普通的 函数 中，则表示跳出该函数，不再执行函数中 return 后面的代码，可以理解成终止函数。
		// 如果 return 语句使用在 main 函数中，表示终止 main 函数，也就是终止程序的运行。
		fmt.Println("无法获取当前路径", err.Error())
		return
	} else {
		fmt.Println(path)
	}
	filePtr, err := os.Open(path + "/ijson/info.json")
	if err != nil {
		// 文件打开失败 open /Users/khmai/awesomego/ijson/info.json: no such file or directory
		fmt.Println("文件打开失败", err.Error())
		return
	}
	defer filePtr.Close()
	var info []Website
	// 创建 json 解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}

	filePtr1, err := os.Open(path + "/ijson/user.json")
	if err != nil {
		// 文件打开失败 open /Users/khmai/awesomego/ijson/info.json: no such file or directory
		fmt.Println("文件打开失败", err.Error())
		return
	}
	defer filePtr1.Close()
	var user []Person
	// 创建 json 解码器
	decoder1 := json.NewDecoder(filePtr1)
	err = decoder1.Decode(&user)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		fmt.Println(user)
	}
}
