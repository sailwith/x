package config

import (
	"os"
	"testing"
)

const config = `
name: hello
age: 10
`

func createTempFile() (string, error) {
	file, err := os.Create("/tmp/config.yaml")
	if err != nil {
		return "", err
	}
	_, err = file.WriteString(config)
	if err != nil {
		return "", err
	}
	if err := file.Close(); err != nil {
		return "", err
	}
	return file.Name(), nil
}

func TestLoad(t *testing.T) {
	filename, err := createTempFile()
	if err != nil {
		t.Error(err)
	}
	t.Log(filename)
	defer os.Remove(filename)

	data := struct {
		Name string
		Age  int
	}{}

	cfg := New()
	if err := cfg.Load(filename); err != nil {
		t.Error(err)
	}
	if err := cfg.Unmarshal(&data); err != nil {
		t.Error(err)
	}
	t.Log(data)
}
