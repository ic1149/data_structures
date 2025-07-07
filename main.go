package main

import (
	"errors"
	"fmt"
)

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

func main() {
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
