package projectcreate

import (
	"flag"
	"fmt"
	"os"
)

var (
	// GitlabToken : Required Gitlab token use to interact with the Gitlab instance
	GitlabToken, _ = os.LookupEnv("GitlabToken")
)

// GenerateCLI : Create new flagset for project's creation CLI
func GenerateCLI() {
	create := flag.NewFlagSet("create", flag.ExitOnError)
	name := create.String("name", "", "Name of the project. Required if path isn't set")
	path := create.String("path", "", "Path of the project. Required if name isn't set")
	description := create.String("description", "", "Description of the project")
	defaultBranch := create.String("default-branch", "", "Set a default branch for the project")
	tags := create.String("tags", "", "List of tags. Coma separated")

	create.Parse(os.Args[3:])
	create.PrintDefaults()
	fmt.Println(*name, *path, *description, *defaultBranch, *tags, GitlabToken)
}
