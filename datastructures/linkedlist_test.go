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
	linkedList.addBack(5)
	linkedList.addBack(10)
	linkedList.addBack(20)
	want := 3
	got := linkedList.size
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
	expected := [3]int{5, 10, 20}
	actual := linkedList.show()
	for i := 0; i < len(actual); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestLinkedListAddFront(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.addBack(5)
	linkedList.addFront(200)
	linkedList.addFront(300)
	want := 3
	got := linkedList.size
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
	expected := [3]int{300, 200, 5}
	actual := linkedList.show()
	for i := 0; i < len(actual); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("expected: %v, got: %v", expected[i], actual[i])
		}
	}
}

func TestLinkedListSearchValue(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.addBack(90)
	linkedList.addFront(119)
	linkedList.addFront(876)
	linkedList.addFront(1000)
	want := true
	got := linkedList.search(876)
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
