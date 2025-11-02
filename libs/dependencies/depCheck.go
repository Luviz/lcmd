package dependencies

import (
	"fmt"
)

type Dependency struct {
	CmdName     string `json:"name" yaml:"name"`
	Description string `json:"description"`
}

type DependencyCheckResult struct {
	Status bool       `json:"status"`
	Dep    Dependency `json:"dependency"`
}

func (dr DependencyCheckResult) String() string {
	statusText := "✔"
	if !dr.Status {
		statusText = "✗"
	}

	return fmt.Sprintf(
		"%s - %s\t%s",
		statusText, dr.Dep.CmdName, dr.Dep.Description)
}

func CheckDependencies() []DependencyCheckResult {
	res := []DependencyCheckResult{}
	for _, dep := range deps {
		res = append(res, CheckDependency(dep))
	}
	return res
}

func CheckDependency(dep Dependency) DependencyCheckResult {
	status := true

	return DependencyCheckResult{
		Status: status,
		Dep:    dep,
	}
}
