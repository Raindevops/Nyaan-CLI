package project

import (
	"fmt"
	"os"
)

var (
	// GitlabToken : Get from os env vars the required token to do http request with the gitlab instance
	GitlabToken, _ = os.LookupEnv("GitlabToken")
	// GitlabURL : Get from os env vars the required gitlab instance's URL
	GitlabURL, _ = os.LookupEnv("GitlabURL")
)

// QueryFields : type of fields use in URL query
type QueryFields struct {
	Name      string
	Queryname string
}

// Project : data type in use in gitlab's Projects
type Project struct {
	Name          QueryFields
	ID            uint
	Description   QueryFields
	Tags          QueryFields
	Path          QueryFields
	Defaultbranch QueryFields
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

func (p *Project) setProjectName(n, qn string) {
	p.Name.Name = n
	p.Name.Queryname = qn
}
func (p *Project) setProjectID(i uint) {
	p.ID = i
}
func (p *Project) setProjectPath(n, qn string) {
	p.Path.Name = n
	p.Path.Queryname = qn
}
func (p *Project) setProjectDesc(n, qn string) {
	p.Description.Name = n
	p.Description.Queryname = qn
}
func (p *Project) setDefaultBranch(n, qn string) {
	p.Defaultbranch.Name = n
	p.Defaultbranch.Queryname = qn
}
func (p *Project) setProjectTags(n, qn string) {
	p.Tags.Name = n
	p.Tags.Queryname = qn
}

func (qf *QueryFields) setName(s string) {
	qf.Name = s
}
func (qf *QueryFields) setQueryName(s string) {
	qf.Queryname = s
}
