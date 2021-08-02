package itemplate

import (
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func DotTemplate() {
	p := Person{"foobar", 18}
	tmpl, _ := template.New("foobar").Parse("Name: {{.Name}}, Age: {{.Age}}")
	// . 表示当前作用域的当前对象，这里表示 p
	// 即运行时 .Name --> p.Name ，.Age --> p.Age
	tmpl.Execute(os.Stdout, p)
}
