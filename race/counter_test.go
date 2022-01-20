package race

import "testing"

func TestCounterSize1(t *testing.T) {
	result := Size("abc")
	if result != 3 {
		t.Error("Incorrect Size")
	}
}

func TestCounterSize2(t *testing.T) {
	result := Size2("abc")
	if result != 3 {
		t.Error("Incorrect Size")
	}
}

func TestCounterSize3(t *testing.T) {
	result := Size3("abc")
	if result != 3 {
		t.Error("Incorrect Size")
	}
}
