package password

import "testing"

func TestEncrypt(t *testing.T) {
	password := "123456"
	hash, err := Encrypt(password)
	if err != nil {
		t.Fatal(err)
	}
	if hash[:3] != "$2a" {
		t.Fatal("encrypt failed", hash, hash[3])
	}
	t.Log(hash)
}

func TestIsValid(t *testing.T) {
	password := "123456"
	hash, err := Encrypt(password)
	if err != nil {
		t.Fatal(err)
	}
	if !IsValid(hash, password) {
		t.Fatal("password error")
	}
	t.Log("password is valid")
}
