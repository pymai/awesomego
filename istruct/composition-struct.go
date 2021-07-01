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

	foobar := Operations{
		Engineer: ops,
		Skill:    skill,
	}

	fmt.Printf("%+v\n", foobar.Engineer)
	fmt.Printf("%+v\n", foobar.Skill)
	fmt.Println(foobar)

}
