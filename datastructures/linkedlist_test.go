package datastructres

import (
	"testing"
)

type test struct {
	input []int
	value int
	want  int
}

func TestLinkedListAddBack(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.AddBack(5)
	linkedList.AddBack(10)
	linkedList.AddBack(20)
	want := 3
	got := linkedList.size
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
	expected := [3]int{5, 10, 20}
	actual := linkedList.Show()
	for i := 0; i < len(actual); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestLinkedListAddFront(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.AddBack(5)
	linkedList.AddFront(200)
	linkedList.AddFront(300)
	want := 3
	got := linkedList.size
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
	expected := [3]int{300, 200, 5}
	actual := linkedList.Show()
	for i := 0; i < len(actual); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestLinkedListSearchValue(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.AddBack(90)
	linkedList.AddFront(119)
	linkedList.AddFront(876)
	linkedList.AddFront(1000)
	want := true
	got := linkedList.Search(876)
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
