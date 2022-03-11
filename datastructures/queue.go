package datastructres

type Queue struct {
	list []int
}

func (q *Queue) Size() int {
	return len(q.list)
}
func (q *Queue) Enqueue(item int) {
	q.list = append(q.list, item)
}

func (q *Queue) Dequeue() int {
	// q.list = append(q.list, item)
	var dequeuedItem int
	dequeuedItem, q.list = q.list[0], q.list[1:]
	return dequeuedItem
}
func (q *Queue) Get() []int {
	return q.list
}
