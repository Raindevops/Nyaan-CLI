package project

import (
	"flag"
	"fmt"
	"os"
)

// CreateCLI : generate the CLI for the create function
func CreateCLI() {
	create := flag.NewFlagSet("create", flag.ExitOnError)
	name := create.String("name", "", "The name for the new project.")

	create.Parse(os.Args[3:])

	fmt.Println(*name)

}
