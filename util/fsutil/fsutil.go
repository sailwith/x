package fsutil

import (
	"os"

	"github.com/gookit/goutil/fsutil"
)

// CreateFile creates a file and automatically creates a directory if the file directory does not exist.
func CreateFile(path string) (*os.File, error) {
	return fsutil.CreateFile(path, 0644, 0755)
}
