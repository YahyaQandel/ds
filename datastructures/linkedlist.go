package datastructres

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	size int
	head *Node
}

type LinkedListI interface {
	addBack()
	show()
}

func (l *LinkedList) addBack(value int) {
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

func (l *LinkedList) show() []int {
	list := []int{}
	for l.head != nil {
		list = append(list, l.head.value)
		l.head = l.head.next
	}
	return list
}

func (l *LinkedList) addFront(value int) {
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

func (l *LinkedList) search(value int) bool {
	list := l.show()
	for i := 0; i < len(list); i++ {
		if list[i] == value {
			return true
		}
	}
	return false
}
