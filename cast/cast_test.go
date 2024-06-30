package cast

import "testing"

func TestToBool(t *testing.T) {
	r := ToBool(1)
	if !r {
		t.Error("cast to bool failed")
	}
}

func TestToString(t *testing.T) {
	r := ToString(123456)
	if r != "123456" {
		t.Error("cast to string failed")
	}
}
