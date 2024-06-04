package double_linked_list

import (
	"fmt"
)

type Node struct {
	Value    int
	NextNode *Node
	PrevNode *Node
}

type OrderedList struct {
	head *Node
	tail *Node
}

func (n Node) Node(value int) {
	n.Value = value
	n.NextNode = nil
	n.PrevNode = nil
}

func initWithValues(values []int, ol *OrderedList) *OrderedList {
	for i := 0; i < len(values); i++ {
		ol.PushBack(values[i])
	}
	return ol
}

func NewList(values ...int) *OrderedList {
	ol := OrderedList{}
	if len(values) > 0 {
		ol = *initWithValues(values, &ol)
	}
	return &ol
}

func (ol *OrderedList) PushBack(value int) {
	newNode := Node{
		Value:    value,
		NextNode: nil,
		PrevNode: nil,
	}
	if ol.head == nil {
		ol.head = &newNode
		ol.tail = &newNode
	} else {
		ol.tail.NextNode = &newNode
		newNode.PrevNode = ol.tail
		ol.tail = &newNode
	}
}

func (ol *OrderedList) PushFront(value int) {
	newNode := Node{
		Value:    value,
		NextNode: nil,
		PrevNode: nil,
	}

	if ol.head == nil {
		ol.head = &newNode
		ol.tail = &newNode
	} else {
		ol.head.PrevNode = &newNode
		newNode.NextNode = ol.head
		ol.head = &newNode
	}
}

func (ol *OrderedList) PopFront() {
	if ol.head == nil {
		panic("list is empty!")
	}

	if ol.head.NextNode == nil {
		ol.head.Value = 0
		ol.head = nil
		ol.tail = nil
		return
	}

	var ptr = ol.head.NextNode
	ol.head.Value = 0
	ol.head.NextNode = nil
	ptr.PrevNode = nil
	ol.head = ptr
}

func (ol *OrderedList) PopBack() {
	if ol.head == nil {
		fmt.Println("List is empty!")
		return
	}

	if ol.head.NextNode == nil {
		ol.head.Value = 0
		ol.head = nil
		ol.tail = nil
		return
	}

	var ptr = ol.tail.PrevNode
	ol.tail.Value = 0
	ol.tail.PrevNode = nil
	ptr.NextNode = nil
	ol.tail = ptr
}

func (ol *OrderedList) FoundElement(position int) int {
	if ol.head == nil {
		panic("list is empty!")
	}

	var lenght int = ol.GetLenght()

	if position+1 > lenght || position*(-1) > lenght {
		panic("going beyond the list!")
	}

	if position < 0 {
		position = position + ol.GetLenght()
	}

	var ptr *Node

	if position <= ol.GetLenght()/2 {
		var count = 0
		ptr = ol.head

		for position != count {
			count++
			ptr = ptr.NextNode
		}
	} else {
		var count = ol.GetLenght() - 1
		ptr = ol.tail

		for position != count {
			count--
			ptr = ptr.PrevNode
		}
	}
	return ptr.Value
}

func (ol *OrderedList) Insert(position int, newValue int) {
	if ol.head == nil {
		panic("list is empty!")
	}

	var lenght int = ol.GetLenght()

	if position+1 > lenght || position*(-1) > lenght+1 {
		panic("going beyond the list!")
	}

	if position < 0 {
		position = position + ol.GetLenght()
	}

	if position == -1 {
		newNode := Node{
			Value:    newValue,
			NextNode: ol.head,
			PrevNode: nil,
		}
		ol.head.PrevNode = &newNode
		ol.head = &newNode
		return
	}

	var ptr *Node

	if position <= ol.GetLenght()/2 {
		var count = 0
		ptr = ol.head

		for position != count {
			count++
			ptr = ptr.NextNode
		}
	} else {
		var count = ol.GetLenght() - 1
		ptr = ol.tail

		for position != count {
			count--
			ptr = ptr.PrevNode
		}
	}

	newNode := Node{
		Value:    newValue,
		NextNode: ptr.NextNode,
		PrevNode: ptr,
	}

	ptr.NextNode = &newNode

	if newNode.NextNode == nil {
		ol.tail = &newNode
		return
	}

	newNode.NextNode.PrevNode = &newNode

	return
}

func (ol *OrderedList) Delete(position int) {
	if ol.head == nil {
		panic("list is empty!")
		return
	}

	var lenght int = ol.GetLenght()

	if position+1 > lenght || position*(-1) > lenght+1 {
		panic("going beyond the list!")
		return
	}

	if ol.head.NextNode == nil {
		ol.head.Value = 0
		ol.head = nil
		ol.tail = nil
		return
	}

	if position < 0 {
		position = position + ol.GetLenght()
	}

	var ptr = ol.head

	if position == 0 {
		ol.head.Value = 0
		ol.head = ol.head.NextNode
		ptr.NextNode = nil
		return
	}

	if position <= ol.GetLenght()/2 {
		var count = 0
		ptr = ol.head

		for position != count {
			count++
			ptr = ptr.NextNode
		}
	} else {
		var count = ol.GetLenght() - 1
		ptr = ol.tail

		for position != count {
			count--
			ptr = ptr.PrevNode
		}
	}

	ptr.Value = 0
	ptr.PrevNode.NextNode = ptr.NextNode

	if ptr == ol.tail {
		ol.tail = ptr.PrevNode
		ptr.PrevNode = nil
		return
	}

	ptr.NextNode.PrevNode = ptr.PrevNode
	ptr.NextNode = nil
	ptr.PrevNode = nil

	return
}

func (ol *OrderedList) GetLenght() int {
	if ol.head == nil {
		panic("list is empty!")
	}

	var count int = 0
	var ptr = ol.head
	for ptr != nil {
		ptr = ptr.NextNode
		count++
	}

	return count
}

func (ol *OrderedList) Show() {
	if ol.head == nil {
		panic("list is empty!")
	}

	var lenght = ol.GetLenght()
	var ptr = ol.head

	for i := 0; i < lenght; i++ {
		if i+1 == lenght {
			fmt.Printf("%d\n", ptr.Value)
		} else {
			fmt.Printf("%d, ", ptr.Value)
		}
		ptr = ptr.NextNode
	}
}

func (ol *OrderedList) Clear() {
	if ol.head == nil {
		panic("list is empty!")
	}

	var ptr = ol.head

	for ptr.NextNode != nil {
		ptr.Value = 0
		ol.head = ptr.NextNode
		ptr.PrevNode = nil
		ptr.NextNode = nil
		ptr = ol.head
	}

	ptr.Value = 0
	ptr.PrevNode = nil
	ol.head = nil
	ol.tail = nil
}

