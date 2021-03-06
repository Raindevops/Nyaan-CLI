package project

import (
	"fmt"
	"os"
)

var (
	// GitlabToken : Get from os env vars the required token to do http request with the gitlab instance
	GitlabToken, _ = os.LookupEnv("GitlabToken")
)

// Project : data type in use in gitlab's Projects
type Project struct {
	Name          string
	ID            uint
	Description   string
	Tags          []string
	Path          string
	Defaultbranch string
}

// RouteProject : entrypoint for the project command line
func RouteProject() {
	fmt.Println("Route project entrypoint for the project command line")

	if len(os.Args) < 3 {
		fmt.Println("project --help invoked")
		return
	}

	switch os.Args[2] {
	case "create":
		CreateCLI()
	case "update":
		fmt.Println("update project invoked")
	case "delete":
		fmt.Println("delete project invoked")
	case "get":
		fmt.Println("get project invoked")
	default:
		fmt.Println("project --help invoked")
		return
	}
}

func (p *Project) setProjectName(s string) {
	p.Name = s
}
func (p *Project) setProjectID(i uint) {
	p.ID = i
}
func (p *Project) setProjectPath(s string) {
	p.Path = s
}
func (p *Project) setProjectDesc(s string) {
	p.Description = s
}
func (p *Project) setDefaultBranch(s string) {
	p.Defaultbranch = s
}
func (p *Project) setProjectTags(s string) {
	p.Tags = append(p.Tags, s)
}
