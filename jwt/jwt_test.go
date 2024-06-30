package jwt

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	claims := map[string]any{
		"id":  1,
		"iss": "issuer",
		"exp": time.Now().Add(time.Hour).Unix(),
		"data": map[string]any{
			"user_id":   1,
			"user_name": "name",
		},
	}
	secret := []byte("123456")
	j := New(secret)
	token, err := j.NewWithClaims(claims)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)

	c, err := j.Parse(token)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(c)
}
