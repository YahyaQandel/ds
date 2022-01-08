package dc_problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfList(t *testing.T) {
	list := []int{1, 3, 4, 5, 43, 20, 93}
	got := sum_of(list)
	want := 169

	assert.Equal(t, want, got)
}

func TestSumOfListWithVariable(t *testing.T) {
	list := []int{1, 3, 4, 5, 43, 20, 93}
	got := sum_of_list_with_variable(list)
	want := 169

	assert.Equal(t, want, got)
}
