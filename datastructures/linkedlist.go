package datastructres

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	size int
	head *Node
}

type ILinkedList interface {
	AddBack()
	AddFront()
	Show()
	Search()
}

func (l *LinkedList) AddBack(value int) {
	newNode := Node{value: value}
	if l.head == nil {
		l.head = &newNode
		l.size += 1
		return
	}
	iterator := l.head
	for iterator.next != nil {
		iterator = iterator.next
	}
	iterator.next = &newNode
	l.size += 1
}

func (l *LinkedList) Show() []int {
	list := []int{}
	for l.head != nil {
		list = append(list, l.head.value)
		l.head = l.head.next
	}
	return list
}

func (l *LinkedList) AddFront(value int) {
	newNode := Node{value: value}
	if l.head == nil {
		l.head = &newNode
		l.size += 1
		return
	}
	temp := l.head
	l.head = &newNode
	l.head.next = temp
	l.size += 1
}

func (l *LinkedList) Search(value int) bool {
	list := l.Show()
	for i := 0; i < len(list); i++ {
		if list[i] == value {
			return true
		}
	}
	return false
}
