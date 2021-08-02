package itemplate

import (
	"net/http"
	"text/template"
)

// TemplateIndex 要有 r *http.Request 参数，不然 itemplate.go 报错
// cannot use itemplate.TemplateIndex (type func(http.ResponseWriter)) as type func(http.ResponseWriter, *http.Request)
// in argument to http.HandleFunc
func TemplateIndex(w http.ResponseWriter, r *http.Request) {
	// 这里的路径是相对于 itemplate.go main() 函数的路径
	info, err := template.ParseFiles("./itemplate/index.html")
	if err != nil {
		panic(err)
	}
	info.Execute(w, "hello world")
}
