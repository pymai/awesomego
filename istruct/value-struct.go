package istruct

import "fmt"

// 结构体是值类型的示例
type WebServer struct {
	Name    string
	Version float64
}

func ValueStruct() {
	var web1 WebServer
	web1.Name = "Nginx"
	web1.Version = 1.1

	// 这里就是值拷贝，所以 web2.Name = "Apache" 的更改，不会影响到 web1.Name
	web2 := web1
	web2.Name = "Apache"

	// 如果想要修改，就要用到指针
	web3 := &web1
	//就是 (*web3).Version = 2.2
	web3.Version = 2.2
	fmt.Println(web1)
	fmt.Println(web2)
	fmt.Println(*web3)
}
