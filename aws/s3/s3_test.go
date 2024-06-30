package s3

import "testing"

func TestPresignPutObject(t *testing.T) {
	s3Client := NewClient("bucket")
	presign, err := s3Client.PresignPutObject("path/to/filename.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(presign.URL)
}

func TestPutObject(t *testing.T) {
	s3Client := NewClient("bucket")
	obj, err := s3Client.PutObject("path/to/filename.txt", []byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(obj.Key, obj.ETag)
}
