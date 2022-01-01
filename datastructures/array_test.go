package datastructres

import (
	"testing"
)

func TestArrayGetIndexByValue(t *testing.T) {
	slice := Slice{}
	slice.Init([]int{1, 2, 3, 4, 5, 6, -90, 1230})
	got := slice.GetIndexByValue(6)
	want := 5
	if want != got {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
	got = slice.GetIndexByValue(100)
	want = -1
	if want != got {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestArrayGetValueByIndex(t *testing.T) {
	slice := Slice{}
	slice.Init([]int{1, 2, 3, 4, 5, 6, -90, 1230})
	got, _ := slice.GetValueByIndex(6)
	want := -90
	if want != got {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
	if want != got {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
func TestArrayGetValueByIndexFailsWithIndexLessThanZero(t *testing.T) {
	slice := Slice{}
	slice.Init([]int{1, 2, 3, 4, 5, 6, -90, 1230})
	got, err := slice.GetValueByIndex(-1)
	want := 0
	wantedErr := "index not valid"
	if err.Error() != wantedErr {
		t.Fatalf("expected to fail with error: %v, got: %v", wantedErr, err.Error())
	}
	if got != 0 {
		t.Fatalf("expected when fails to return: %v, got: %v", want, got)
	}
}

func TestArrayGetValueByIndexFailsWithIndexOutOfRange(t *testing.T) {
	slice := Slice{}
	slice.Init([]int{1, 2, 3, 4, 5, 6, -90, 1230})
	got, err := slice.GetValueByIndex(99)
	want := 0
	wantedErr := "index out of range"
	if err.Error() != wantedErr {
		t.Fatalf("expected to fail with error: %v, got: %v", wantedErr, err.Error())
	}
	if got != 0 {
		t.Fatalf("expected when fails to return: %v, got: %v", want, got)
	}
}
