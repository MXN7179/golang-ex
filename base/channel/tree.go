package main

type binary struct {
	v     interface{}
	left  *binary
	right *binary
}

func main() {
	height()
}

func height(node *binary) uint {
	if node == nil {
		return 0
	}
	if node.right == nil || node.left == nil {
		return 1
	}
	res := 1 + _max(height(node.left), height(node.right))
	return res
}

func _max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

type list struct {
	value interface{}
	pre   *list
	next  *list
}

func reverse(l *list) *list {
	if l == nil {
		return nil
	}
	head := l // head
	p := head
	for p.next != nil {
		p = p.next
	}
	tail := p // tail

	for head != tail && head.next != tail { // even  odd
		tmp := head.value
		head.value = tail.value
		tail.value = tmp

		head = head.next
		tail = tail.pre
	}
	return l
}
