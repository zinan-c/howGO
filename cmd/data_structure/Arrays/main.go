package main

import (
	"container/list"
	"fmt"
	"howgo/internal/model"
)

func main() {
	//normal()

	//listNode()

	//linkedList()
}

func linkedList() {
	list := model.NewLinkedList()
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(3)
	list.AddFirst(0)
	list.Add(2, 100)

	list.Display()
	// size = 5
	// 0 <-> 1 <-> 100 <-> 2 <-> 3 <-> null
}
func listNode() {
	head := model.CreateLinkedList([]int{1, 2, 3, 4, 5})

	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}

	// insert to first element
	newNode := &model.ListNode{Val: 0}
	newNode.Next = head
	head = newNode

	// 0 -> 1 -> 2 -> 3 -> 4 -> 5
	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}

	// in the tail insert element 6
	p := head
	// goto latest
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &model.ListNode{Val: 6}

	// 0 -> 1 -> 2 -> 3 -> 4 -> 5 ->6
	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}

	headDoubly := model.CreateDoublyLinkedList([]int{1, 2, 3, 4, 5})
	var tail *model.DoublyListNode

	for p := headDoubly; p != nil; p = p.Next {
		fmt.Println(p.Val)
		tail = p
	}
	for p := tail; p != nil; p = p.Prev {
		fmt.Println(p.Val)
	}
}

func normal() {
	/**
		** List
	**/
	// initial the list
	lst := list.New()
	lst.PushBack(1)
	lst.PushBack(2)
	lst.PushBack(3)
	lst.PushBack(4)
	lst.PushBack(5)

	fmt.Println(lst.Len() == 0)
	fmt.Println(lst.Len())

	// put one element in head
	lst.PushFront(0)
	// put one element in tail
	lst.PushBack(6)

	// get the head element and tail element
	front := lst.Front().Value.(int)
	back := lst.Back().Value.(int)
	fmt.Println(front, back)

	// delete the head element
	lst.Remove(lst.Front())
	// delete the tail element
	lst.Remove(lst.Back())

	// insert element in list
	// move to third position
	third := lst.Front().Next().Next()
	lst.InsertBefore(99, third)

	// delete some element in this list
	second := lst.Front().Next()
	lst.Remove(second)

	// iterate list
	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value.(int), " ")
	}
	fmt.Println()

	/**
		** Queue
	**/
	// initial the q with list
	q := list.New()

	// add element in tail
	q.PushBack(10)
	q.PushBack(20)
	q.PushBack(30)

	fmt.Println(q.Len() == 0)
	fmt.Println(q.Len())

	// get the head
	front = q.Front().Value.(int)
	fmt.Println(front)

	// delete the head element
	q.Remove(q.Front())

	// get the new head
	newFront := q.Front().Value.(int)
	fmt.Println(newFront)

	/**
		** Stack
	**/
	// initial one int stack
	var s []int

	// put element into stack head(slice tail)
	s = append(s, 10)
	s = append(s, 20)
	s = append(s, 30)

	fmt.Println(len(s) == 0)
	fmt.Println(len(s))

	// get the stack head element
	fmt.Println(s[len(s)-1])

	// delete current stack head element
	s = s[:len(s)-1]

	// get the new head
	fmt.Println(s[len(s)-1])

	/**
		** HashMap
	**/
	// initial one empty hashmap
	// var hashmap = make(map[int]string)

	// // initial one hashmap with value
	// hashmap = map[int]string{
	//     1: "one",
	//     2: "two",
	//     3: "three",
	// }
	// initial empty then put the value
	hmap := make(map[int]string)
	hmap[1] = "one"
	hmap[2] = "two"
	hmap[3] = "three"

	fmt.Println(len(hmap) == 0)
	fmt.Println(len(hmap))

	// check some element exist or not
	// output：Key 2 -> two
	if val, exists := hmap[2]; exists {
		fmt.Println("Key 2 ->", val)
	} else {
		fmt.Println("Key 2 not found.")
	}

	// get value of current key, if does not exist, return empty string
	// 输出：
	fmt.Println(hmap[4])

	hmap[4] = "four"
	fmt.Println(hmap[4])

	// delete which map
	delete(hmap, 3)

	// output：Key 3 not found.
	if val, exists := hmap[3]; exists {
		fmt.Println("Key 3 ->", val)
	} else {
		fmt.Println("Key 3 not found.")
	}

	// output (may sorted)：
	// 1 -> one
	// 2 -> two
	// 4 -> four
	for key, value := range hmap {
		fmt.Printf("%d -> %s\n", key, value)
	}

	/**
		** HashSet
	**/
	// initial one hashset with some values
	hashset := map[int]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
	}

	fmt.Println(len(hashset) == 0)
	fmt.Println(len(hashset))

	// output：Element 3 found.
	if _, exists := hashset[3]; exists {
		fmt.Println("Element 3 found.")
	} else {
		fmt.Println("Element 3 not found.")
	}

	hashset[5] = struct{}{}

	delete(hashset, 2)
	// output：Element 2 not found.
	if _, exists := hashset[2]; exists {
		fmt.Println("Element 2 found.")
	} else {
		fmt.Println("Element 2 not found.")
	}

	// output (may sorted)：
	// 1
	// 3
	// 4
	// 5
	for element := range hashset {
		fmt.Println(element)
	}
}
