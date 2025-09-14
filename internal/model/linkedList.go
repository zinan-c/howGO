package model

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewLinkedList() *LinkedList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &LinkedList{head: head, tail: tail, size: 0}
}

func (list *LinkedList) AddLast(e interface{}) {
	x := &Node{val: e}
	temp := list.tail.prev
	// temp <-> x
	temp.next = x
	x.prev = temp

	x.next = list.tail
	list.tail.prev = x
	// temp <-> x <-> tail
	list.size++
}

func (list *LinkedList) AddFirst(e interface{}) {
	x := &Node{val: e}
	temp := list.head.next
	// head <-> temp
	temp.prev = x
	x.next = temp

	list.head.next = x
	x.prev = list.head
	// head <-> x <-> temp
	list.size++
}

func (list *LinkedList) Add(index int, element interface{}) error {
	if err := list.checkPositionIndex(index); err != nil {
		return err
	}
	if index == list.size {
		list.AddLast(element)
		return nil
	}

	p := list.getNode(index)
	temp := p.prev
	// temp <-> p
	x := &Node{val: element}

	p.prev = x
	temp.next = x

	x.prev = temp
	x.next = p
	// temp <-> x <-> p
	list.size++
	return nil
}

func (list *LinkedList) RemoveFirst() (interface{}, error) {
	if list.size < 1 {
		return nil, errors.New("no elements to remove")
	}
	// 虚拟节点的存在是我们不用考虑空指针的问题
	x := list.head.next
	temp := x.next
	// head <-> x <-> temp
	list.head.next = temp
	temp.prev = list.head

	// head <-> temp

	list.size--
	return x.val, nil
}

func (list *LinkedList) RemoveLast() (interface{}, error) {
	if list.size < 1 {
		return nil, errors.New("no elements to remove")
	}
	x := list.tail.prev
	temp := x.prev
	// temp <-> x <-> tail

	list.tail.prev = temp
	temp.next = list.tail

	// temp <-> tail

	list.size--
	return x.val, nil
}

func (list *LinkedList) Remove(index int) (interface{}, error) {
	if err := list.checkElementIndex(index); err != nil {
		return nil, err
	}
	// 找到 index 对应的 Node
	x := list.getNode(index)
	prev := x.prev
	next := x.next
	// prev <-> x <-> next
	prev.next = next
	next.prev = prev

	list.size--

	return x.val, nil
}

func (list *LinkedList) Get(index int) (interface{}, error) {
	if err := list.checkElementIndex(index); err != nil {
		return nil, err
	}
	p := list.getNode(index)

	return p.val, nil
}

func (list *LinkedList) GetFirst() (interface{}, error) {
	if list.size < 1 {
		return nil, errors.New("no elements in the list")
	}

	return list.head.next.val, nil
}

func (list *LinkedList) GetLast() (interface{}, error) {
	if list.size < 1 {
		return nil, errors.New("no elements in the list")
	}

	return list.tail.prev.val, nil
}

func (list *LinkedList) Set(index int, val interface{}) (interface{}, error) {
	if err := list.checkElementIndex(index); err != nil {
		return nil, err
	}
	p := list.getNode(index)

	oldVal := p.val
	p.val = val

	return oldVal, nil
}

func (list *LinkedList) Size() int {
	return list.size
}

func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *LinkedList) getNode(index int) *Node {
	p := list.head.next
	for i := 0; i < index; i++ {
		p = p.next
	}
	return p
}

func (list *LinkedList) isElementIndex(index int) bool {
	return index >= 0 && index < list.size
}

func (list *LinkedList) isPositionIndex(index int) bool {
	return index >= 0 && index <= list.size
}

func (list *LinkedList) checkElementIndex(index int) error {
	if !list.isElementIndex(index) {
		return fmt.Errorf("Index: %d, Size: %d", index, list.size)
	}
	return nil
}

func (list *LinkedList) checkPositionIndex(index int) error {
	if !list.isPositionIndex(index) {
		return fmt.Errorf("Index: %d, Size: %d", index, list.size)
	}
	return nil
}

func (list *LinkedList) Display() {
	fmt.Printf("size = %d\n", list.size)
	p := list.head.next
	for p != list.tail {
		fmt.Printf("%v <-> ", p.val)
		p = p.next
	}
	fmt.Println("null\n")
}
