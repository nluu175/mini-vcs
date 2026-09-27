package commands

import (
	"flag"
	"fmt"
	"os"

	"mini-vcs/internal/objects"
)

// RunHashObject computes the blob hash for a file's content, optionally
// writing it to the object store.
func RunHashObject(args []string) error {
	fs := flag.NewFlagSet("hash-object", flag.ExitOnError)
	write := fs.Bool("w", false, "write the object to the store")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if fs.NArg() != 1 {
		return fmt.Errorf("usage: vcs hash-object [-w] <file>")
	}
	path := fs.Arg(0)

	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading %s: %w", path, err)
	}

	hash, data := objects.HashBlob(content)

	if *write {
		if err := objects.WriteObject(hash, data); err != nil {
			return err
		}
	}

	fmt.Println(hash)
	return nil
}
