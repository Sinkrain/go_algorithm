package reverse

import (
	"fmt"
	. "linkedList/deduplicate"
)


// ReverseLinkedList reversee linkedList
func ReverseLinkedList(n int){
	CenterMsg("Reverse LinkedList")
	numSlice := BuildNumSlice(n, 100)
	fmt.Println("numSlice: ", numSlice)
	l := BuildLinkedList(numSlice)
	ReadLinkedSlice(&l)
	Reverse(&l)
	ReadLinkedSlice(&l)
}

// Reverse reverse LinkedList
func Reverse(lnode *LNode){
	if lnode == nil || lnode.Next == nil{
		return
	}
	pre := &LNode{}
	cur := &LNode{}
	next := lnode.Next
	for next != nil{
		cur = next.Next
		if pre.Next == nil && pre.Num ==0{
			next.Next = nil
		} else {
			next.Next = pre
		}
		pre = next
		next = cur
	}
	lnode.Next = pre
}