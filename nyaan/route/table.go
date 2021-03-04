package route

import (
	"fmt"
	"os"

	"gitlab.com/raindevops/nyaan-cli/nyaan/entities/project"
)

var (
	// GitlabToken : Required Gitlab token use to interact with the Gitlab instance
	GitlabToken, _ = os.LookupEnv("gitlab-token")
)

// Table : route table for cli
func Table() {
	if len(os.Args) < 2 {
		fmt.Println("--help command invokes")
		return
	}

	switch os.Args[1] {
	case "project":
		project.RouteProject()
	default:
		fmt.Println("--help command invoked")
		return
	}
}
