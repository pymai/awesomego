### go 项目初始化
```
// 初始化当前文件夹，创建 go.mod 文件
$ go mod init awesomego
// 增加缺少的包，删除无用的包
$ go mod tidy
// go.mod 的 module 要以域名开头，不然会报错 "missing dot in first path element"
module github.com/pymai/awesomego
```

