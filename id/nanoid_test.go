package id

import "testing"

func TestNewNanoID(t *testing.T) {
	id, err := NewNanoID(10)
	if err != nil {
		t.Fatal(err)
	}
	if len(id) != 10 {
		t.Fatal("id length is not 10", id)
	}
	t.Log("id length is 10", id)

	id2, err := NewNanoID(20)
	if err != nil {
		t.Fatal(err)
	}
	if len(id2) != 20 {
		t.Fatal("id length is not 20", id2)
	}
	t.Log("id length is 20", id2)
}
