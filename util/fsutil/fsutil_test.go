package fsutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFile(t *testing.T) {
	path := "/tmp/tmpfile"
	f, err := CreateFile(path)
	assert.NoError(t, err)
	defer os.RemoveAll(path)
	t.Log(f.Name())
}
