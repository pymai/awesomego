package istruct

import "fmt"

type Operations struct {
	Engineer SystemOperationEngineer
	Skill    SystemOperationSkill
}

type Sre struct {
	// struct 嵌入
	// 这样就可以直接调用 SystemOperationSkill 的 rm() 方法
	// 也可以直接访问 SystemOperationEngineer 的 Name 成员
	SystemOperationEngineer
	SystemOperationSkill
}

type SystemOperationEngineer struct {
	Name          string
	Age           int64
	WorkSeniority int64
}

type SystemOperationSkill struct {
	ProfessionalSkill   string
	EducationBackground string
}

func (s SystemOperationSkill) rm() string {
	return "rm -rf"
}

// 结构体接口转发
func (o Operations) rm() string {
	return o.Skill.rm()
}

func CompositionStruct() {
	ops := SystemOperationEngineer{
		Name:          "foo",
		Age:           25,
		WorkSeniority: 1,
	}
	skill := SystemOperationSkill{
		ProfessionalSkill:   "Linux Golang",
		EducationBackground: "University Degree",
	}
	// 结构体初始化后，就可以调用结构体的方法
	null := skill.rm()
	fmt.Println(null)

	foobar := Operations{
		Engineer: ops,
		Skill:    skill,
	}

	fmt.Printf("%+v\n", foobar.Engineer)
	fmt.Printf("%+v\n", foobar.Skill)
	fmt.Println(foobar)
	fmt.Println(foobar.rm())

	// 通过 结构体嵌入 直接调用 "子结构体" 的方法
	foo := Sre{ops, skill}
	fmt.Println(foo.rm())
	fmt.Println(foo.Name)
}
