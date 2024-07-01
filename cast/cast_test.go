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

func TestToInt(t *testing.T) {
	r := ToInt("-123456")
	if r != -123456 {
		t.Error("cast to int failed")
	}
}

func TestToUInt(t *testing.T) {
	r := ToUInt("123456")
	if r != 123456 {
		t.Error("cast to uint failed")
	}
}

func TestToInt64(t *testing.T) {
	r := ToInt64("-1444784865584")
	if r != -1444784865584 {
		t.Error("cast to int64 failed")
	}
}

func TestToUInt64(t *testing.T) {
	r := ToUInt64("1444784865584")
	if r != 1444784865584 {
		t.Error("cast to uint64 failed")
	}
}
