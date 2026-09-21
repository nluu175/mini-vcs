package main

import (
	"fmt"
	"os"

	"mini-vcs/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: vcs <command> [<args>]")
		os.Exit(1)
	}

	var err error
	switch os.Args[1] {
	case "init":
		err = commands.RunInit(os.Args[2:])
	default:
		err = fmt.Errorf("unknown command: %s", os.Args[1])
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
