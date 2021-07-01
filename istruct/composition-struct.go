package istruct

import "fmt"

type Operations struct {
	Engineer SystemOperationEngineer
	Skill    SystemOperationSkill
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

}
