package objects

import (
	"fmt"
	"os"
	"path/filepath"
)

const objectsDir = ".vcs/objects"

// objectPath is the one place that knows how a hash maps to a file path
// (flat layout — see DESIGN.md D5). Everything else should call this
// rather than build the path itself.
func objectPath(hash string) string {
	return filepath.Join(objectsDir, hash)
}

// WriteObject writes raw object data (header+content) to the store under
// its hash. Content-addressed, so if it's already there, there's nothing
// to do.
func WriteObject(hash string, data []byte) error {
	path := objectPath(hash)
	if _, err := os.Stat(path); err == nil {
		return nil
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("writing object %s: %w", hash, err)
	}
	return nil
}
