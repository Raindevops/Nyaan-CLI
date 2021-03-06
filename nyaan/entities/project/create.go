package project

import (
	"flag"
	"fmt"
	"os"
)

var project Project

// CreateCLI : generate the CLI for the create function
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

	fmt.Println(*name, *desc, *tags, *path, *defBranch)
	create.PrintDefaults()

}
