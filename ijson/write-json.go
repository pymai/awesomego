package ijson

import (
	"encoding/json"
	"fmt"
	"os"
)

type Website struct {
	Name   string
	Url    string
	Course []string
}

type Person struct {
	Name string
	Age  int32
}

func WriteJson() {
	// 一个 slice ，类型是 Website 结构体
	// 2021.05.20 如果 json 里面既有字符串又有数字，结构体要怎样定义？
	// 2021.06.11 如果 json 里面既有字符串又有数字，结构体要怎样定义？--> 结构体里字符串定义为 string，数字定义为 int32
	info := []Website{
		{
			"Golang",
			"http://c.biancheng.net/golang/",
			[]string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"},
		},
		{
			"java",
			"http://c.biancheng.net/java/",
			[]string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"},
		},
	}

	user := []Person{
		{
			Name: "foo",
			Age:  20,
		},
		{
			Name: "bar",
			Age:  30,
		},
	}

	// 创建文件
	path, err := os.Getwd()
	if err != nil {
		// 如果 return 语句使用在普通的 函数 中，则表示跳出该函数，不再执行函数中 return 后面的代码，可以理解成终止函数。
		// 如果 return 语句使用在 main 函数中，表示终止 main 函数，也就是终止程序的运行。
		fmt.Println("无法获取当前路径", err.Error())
		return
	} else {
		fmt.Println(path)
	}
	filePtr, err := os.Create(path + "/ijson/info.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()

	// 创建 json 编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())
	} else {
		fmt.Println("编码成功")
	}

	// 创建文件
	path1, err := os.Getwd()
	if err != nil {
		// 如果 return 语句使用在普通的 函数 中，则表示跳出该函数，不再执行函数中 return 后面的代码，可以理解成终止函数。
		// 如果 return 语句使用在 main 函数中，表示终止 main 函数，也就是终止程序的运行。
		fmt.Println("无法获取当前路径", err.Error())
		return
	} else {
		fmt.Println(path1)
	}
	filePtr1, err := os.Create(path + "/ijson/user.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr1.Close()

	// 创建 json 编码器
	encoder1 := json.NewEncoder(filePtr1)
	err = encoder1.Encode(user)
	if err != nil {
		fmt.Println("编码错误", err.Error())
	} else {
		fmt.Println("编码成功")
	}
}
