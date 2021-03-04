package route

import "flag"

// Table : route table for cli
func Table() {
	create := flag.NewFlagSet("create", flag.ExitOnError)
	typePtr := create.String("type", "", "The entity you want to create. project, issue...")
	projectId := create.String("projet-id", "", "The unic identifier for the project ")
}
