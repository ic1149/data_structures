package main

import (
	"errors"
	"fmt"
	"slices"
)

type graph_connection struct {
	node   graph_node
	weight int
}
type graph_node struct {
	to []graph_connection
}

type tree_root struct {
	name       string
	first_node *tree_node
}

type tree_node struct {
	data any
	l, r *tree_node
}

type stack struct {
	data []any
}

func (st *stack) peek() any {
	if len(st.data) > 0 {
		return st.data[len(st.data)-1]
	} else {
		return nil
	}
}

func (st *stack) pop() any {
	if len(st.data) > 0 {
		grab := st.data[len(st.data)-1]
		st.data = slices.Delete(st.data, len(st.data)-1, 1)
		return grab
	} else {
		return nil
	}
}

func (st *stack) push(data any) {
	st.data = append(st.data, data)
}

type queue struct {
	data []any
}

func (q *queue) peek() any {
	if len(q.data) > 0 {
		return q.data[0]
	} else {
		return nil
	}
}

func (q *queue) deq() any {
	if len(q.data) > 0 {
		grab := q.data[0]
		q.data = slices.Delete(q.data, 0, 1)
		return grab
	} else {
		return nil
	}
}

func (q *queue) enq(data any) {
	q.data = append(q.data, data)
}

type headPointer *node

type linkedList struct {
	name  string
	headP headPointer
}

type node struct {
	data            any
	pointer_to_next *node
}

func (l *linkedList) output() {
	var outstr string
	var this node
	next := l.headP
	for next != nil {
		this = *next // data at pointer next
		next = this.pointer_to_next
		fmt.Println(this.data)
		outstr += fmt.Sprintf("%v, ", this.data)
	}
	outstr = outstr[:len(outstr)-2]
	fmt.Printf("Linked List %s: {%v}\n", l.name, outstr)
}

func (l *linkedList) exists(target any) bool {
	var this node
	next := l.headP
	for next != nil {
		this = *next // data at pointer next
		if target == this.data {
			return true
		}
		next = this.pointer_to_next
	}

	return false
}

func (l *linkedList) append(newdata any) {
	var this node

	next := l.headP
	for {
		this = *next // data at pointer next
		if this.pointer_to_next != nil {
			next = this.pointer_to_next
		} else {
			break
		}
	}

	newnode := node{data: newdata, pointer_to_next: nil}
	*next = node{data: this.data, pointer_to_next: &newnode}
	// set pointer to next, on original last node, to the new node
}

func (l *linkedList) delete(target any) error {
	var this node
	var peek_next node
	next := l.headP
	for next != nil {
		this = *next // data at pointer next
		peek_next = *this.pointer_to_next
		if target == peek_next.data {
			*next = node{data: this.data, pointer_to_next: peek_next.pointer_to_next}
			// if this node is to be removed, set the pointer to next, on the last node, to the next node
			return nil
		}
		next = this.pointer_to_next
	}

	return errors.New("target not found")
}

func ll_test() {
	c := node{data: "apple", pointer_to_next: nil}
	b := node{data: "banana", pointer_to_next: &c}
	a := node{data: "orange", pointer_to_next: &b}
	var fruit = linkedList{name: "fruit", headP: &a}
	fruit.output()
	fmt.Println("removing apple")
	err := fruit.delete("apple")
	if err != nil {
		fmt.Println(err)
		return
	}
	if fruit.exists("apple") {
		fmt.Println("apple is in the linked list")
	} else {
		fmt.Println("apple is not in the linked list")
	}
	fruit.output()
	fmt.Println("appending watermelon")
	fruit.append("watermelon")
	fruit.output()
}

func q_test() {
	var q1 queue
	q1.enq("breakfast")
	q1.enq("lunch")
	q1.enq("tea")
	fmt.Println(len(q1.data))
	fmt.Println(q1.peek())
	fmt.Println(q1.deq())
	fmt.Println(len(q1.data))
}

func main() {
	fmt.Println("Welcome to Data Structures")
	q_test()
}
