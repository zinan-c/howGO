package main

import (
    "container/list"
    "fmt"
)

func main() {
	/**
		** List
	**/
    // initial the list
    list := list.New()
    list.PushBack(1)
    list.PushBack(2)
    list.PushBack(3)
    list.PushBack(4)
    list.PushBack(5)

    fmt.Println(list.Len() == 0)
    fmt.Println(list.Len())

    // put one element in head
    list.PushFront(0)
    // put one element in tail
    list.PushBack(6)

    // get the head element and tail element
    front := list.Front().Value.(int)
    back := list.Back().Value.(int)
    fmt.Println(front, back)

    // delete the head element
    list.Remove(list.Front())
    // delete the tail element
    list.Remove(list.Back())

    // insert element in list
    // move to third position
    third := list.Front().Next().Next()
    list.InsertBefore(99, third)

    // delete some element in this list
    second := list.Front().Next()
    list.Remove(second)

    // iterate list
    for e := list.Front(); e != nil; e = e.Next() {
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
    front := q.Front().Value.(int)
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
    var hashmap map[int]string
    hashmap = make(map[int]string)

    // initial one hashmap with value
    hashmap = map[int]string{
        1: "one",
        2: "two",
        3: "three",
    }
	// initial empty then put the value
    hashmap := make(map[int]string)
    hashmap[1] = "one"
    hashmap[2] = "two"
    hashmap[3] = "three"

    fmt.Println(len(hashmap) == 0)
    fmt.Println(len(hashmap))

    // check some element exist or not
    // output：Key 2 -> two
    if val, exists := hashmap[2]; exists {
        fmt.Println("Key 2 ->", val)
    } else {
        fmt.Println("Key 2 not found.")
    }

    // get value of current key, if does not exist, return empty string
    // 输出：
    fmt.Println(hashmap[4])

    hashmap[4] = "four"
    fmt.Println(hashmap[4])

    // delete which map
    delete(hashmap, 3)

    // output：Key 3 not found.
    if val, exists := hashmap[3]; exists {
        fmt.Println("Key 3 ->", val)
    } else {
        fmt.Println("Key 3 not found.")
    }

    // output (may sorted)：
    // 1 -> one
    // 2 -> two
    // 4 -> four
    for key, value := range hashmap {
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