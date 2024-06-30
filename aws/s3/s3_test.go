package s3

import "testing"

func TestPresignPutObject(t *testing.T) {
	client := NewClient("bucket")
	presign, err := client.PresignPutObject("path/to/filename.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(presign.URL)
}

func TestPutObject(t *testing.T) {
	client := NewClient("bucket")
	obj, err := client.PutObject("path/to/filename.txt", []byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(obj.Key, obj.ETag)
}
