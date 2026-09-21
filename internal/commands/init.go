package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

const vcsDir = ".vcs"

// RunInit creates the object store directory structure for a new repo.
// It errors if a repo already exists in the current directory.
func RunInit(args []string) error {
	if _, err := os.Stat(vcsDir); err == nil {
		return fmt.Errorf("%s already exists", vcsDir)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("checking %s: %w", vcsDir, err)
	}

	objectsDir := filepath.Join(vcsDir, "objects")
	if err := os.MkdirAll(objectsDir, 0755); err != nil {
		return fmt.Errorf("creating %s: %w", objectsDir, err)
	}

	fmt.Printf("Initialized empty vcs repository in %s\n", vcsDir)
	return nil
}
