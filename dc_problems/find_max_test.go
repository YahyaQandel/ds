package dc_problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMax(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 100, 43, 594, 20}
	got := findMax(sample)
	want := 594

	assert.Equal(t, want, got)
}
