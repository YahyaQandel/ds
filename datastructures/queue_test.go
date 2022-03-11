package datastructres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueueItem(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(10)
	assert.Equal(t, queue.Get(), []int{10})
}
func TestEnqueueMultipleItems(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(-3)
	queue.Enqueue(200)
	queue.Enqueue(14)
	assert.Equal(t, queue.Get(), []int{10, -3, 200, 14})
}
func TestDequeueItem(t *testing.T) {
	queue := Queue{
		list: []int{14},
	}
	dequeuedItem := queue.Dequeue()
	assert.Equal(t, queue.Get(), []int{})
	assert.Equal(t, dequeuedItem, 14)
}

func TestDequeueMultipleItems(t *testing.T) {
	queue := Queue{
		list: []int{14, 9, -2, 0},
	}
	dequeuedItem := queue.Dequeue()
	assert.Equal(t, queue.Get(), []int{9, -2, 0})
	assert.Equal(t, dequeuedItem, 14)
}
