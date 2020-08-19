package sum

import (
	"fmt"
	. "linkedList/deduplicate"
)


// SumLinkedList sum for linkedList
func SumLinkedList(n int){
	fmt.Println("\n\nSumLinkedList")

	l := createLinkedList(n)
	r := createLinkedList(n)

	// SumLinkedListByCount(l, r)

	SumLinkedListByAdd(l, r)
}


func createLinkedList(n int)LNode{
	
	numSlice := BuildNumSlicet(n, 10)
	fmt.Println("Num Slice: ", numSlice)
	numLinkedList := BuildLinkedList(numSlice)
	ReadLinkedSlice(&numLinkedList)
	return numLinkedList
}