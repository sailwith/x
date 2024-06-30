package id

import "testing"

func TestNewRandomNumber(t *testing.T) {
	num, err := NewRandomNumber(6)
	if err != nil {
		t.Error("random number error:", err)
	}
	if num < 100000 || num > 999999 {
		t.Error("random number not in range:", num)
	}
}
