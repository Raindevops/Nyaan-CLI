package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Helper command is called")
		return
	}

	fmt.Println("Entrypoint")
}
