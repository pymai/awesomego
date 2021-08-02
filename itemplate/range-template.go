package itemplate

import (
	"os"
	"text/template"
)

type Tool struct {
	Name string
}

type Worker struct {
	UserName string
	Email    []string
	Tool     []*Tool
}

func RangeTemplate() {
	t1 := Tool{Name: "锤子"}
	t2 := Tool{Name: "钉子"}
	t := template.New("range-template")
	// 2021.07.29
	// 1. 为什么 Email 直接 range，而 Tool 要先 with 再 range ?
	// 2021.08.03 --> Worker 的 Tool 指向另一层结构体 Tool
	// 2. - 是什么意思 ？
	// 2021.08.03 {{- --> 去除前面的空白行; - }} --> 去除后面的空白行
	// 3. 为什么 - 一定要和 [[ 连接在一起？如果有中间有空格（[[ -）,就会报错 illegal number syntax: "-"
	// 2021.08.03 --> 语法是这样
	// 4. {{range}} {{end}}
	// 5. {{with}} {{end}}
	t = template.Must(t.Parse(
		`hello {{ .UserName }}!
{{- range .Email }}
your email is {{ . }}
{{-  end -}}
{{ with .Tool }}
{{-  range . }}
your tool is {{ .Name }}
{{- end -}}
{{ end }}
{{"output" | printf "%q" -}}
{{ $len := (len "foobar")}}
{{ println $len -}}
`))
	p := Worker{
		UserName: "foobar",
		Email:    []string{"admin@qq.com", "admin@gmail.com"},
		Tool:     []*Tool{&t1, &t2},
	}
	t.Execute(os.Stdout, p)
}
