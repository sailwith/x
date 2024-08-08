package s3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPresignPutObject(t *testing.T) {
	client := NewClient("bucket")
	presign, err := client.PresignPutObject("path/to/filename.txt")
	assert.NoError(t, err)
	t.Log(presign.URL)
}

func TestPutObject(t *testing.T) {
	client := NewClient("bucket")
	obj, err := client.PutObject("path/to/filename.txt", []byte("hello world"))
	assert.NoError(t, err)
	t.Log(obj.Key, obj.ETag)
}
