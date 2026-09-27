package objects

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

// HashBlob computes the object hash for blob content using git's
// "blob <size>\0<content>" header format. It returns the hex-encoded
// hash and the full header+content bytes, so a caller that wants to
// write the object doesn't need to rebuild them.
func HashBlob(content []byte) (hash string, data []byte) {
	header := fmt.Sprintf("blob %d\x00", len(content))
	data = append([]byte(header), content...)
	sum := sha1.Sum(data)
	return hex.EncodeToString(sum[:]), data
}
