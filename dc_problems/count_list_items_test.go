package dc_problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountItemsInList(t *testing.T) {
	list := []int{1, 3, 4, 5, 43, 20, 93, 4, 5, 43, 20, 93, 4, 5, 43, 20, 93, 4, 5, 43, 20, 93}
	got := count_of(list)
	want := 22
	assert.Equal(t, want, got)
	list = []int{}
	got = count_of(list)
	want = 0
	assert.Equal(t, want, got)
}
