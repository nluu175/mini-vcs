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

	cmds := map[string]func([]string) error{
		"init":        commands.RunInit,
		"hash-object": commands.RunHashObject,
	}

	run, ok := cmds[os.Args[1]]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}

	if err := run(os.Args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
