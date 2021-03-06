package project

import (
	"fmt"
	"os"
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
