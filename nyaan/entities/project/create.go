package project

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

var project Project

// CreateCLI : generate the CLI for the create function
// TODO : refacto flag declaration & storage in structure
func CreateCLI() {
	create := flag.NewFlagSet("create", flag.ExitOnError)
	name := create.String("name", "", "The name for the new project. (Required if no path set)")
	desc := create.String("description", "", "The description for the new project. (Optional)")
	tags := create.String("tags", "", "List of tags for the new project. They need to be coma separated.")
	path := create.String("path", "", "The path for the new project. (Required if no name set)")
	defBranch := create.String("default-branch", "master", "The default branch for the new project. (Optional)")

	create.Parse(os.Args[3:])

	if *name == "" && *path == "" {
		create.PrintDefaults()
		return
	}

	if *name != "" {
		project.setProjectName(*name)
	}
	if *desc != "" {
		project.setProjectDesc(*desc)
	}
	if *path != "" {
		project.setProjectPath(*path)
	}
	if *defBranch != "" {
		project.setDefaultBranch(*defBranch)
	}
	if *tags != "" {
		ar := strings.Split(*tags, ",")
		for _, s := range ar {
			project.setProjectTags(s)
		}
	}
	CreateProject()
}

// CreateProject : Do an http request with the Project structure informations
func CreateProject() {
	f := reflect.ValueOf(project)

	for i := 0; i < f.NumField(); i++ {
		if !f.Field(i).IsZero() {
			fmt.Println(f.Field(i).Interface())
		}
	}
}
