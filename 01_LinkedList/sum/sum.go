package sum

import (
	"fmt"
	. "linkedList/deduplicate"
)


// SumLinkedList sum for linkedList
func SumLinkedList(n int){

	l := createLinkedList(n)
	r := createLinkedList(n+2)

	SumLinkedListByCount(l, r)
	SumLinkedListByAdd(l, r)
}


func createLinkedList(n int)LNode{
	
	numSlice := BuildNumSlice(n, 10)
	fmt.Println("Num Slice: ", numSlice)
	numLinkedList := BuildLinkedList(numSlice)
	ReadLinkedSlice(&numLinkedList)
	return numLinkedList
}