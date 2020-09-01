package flip

import (
	// "fmt"
	. "linkedList/deduplicate"
)

// FlipLinkedList Flip linkedList
func FlipLinkedList(n int){
	CenterMsg("Flip LinkedList")
	numSlice := BuildNumSlice(n, 100)
	lnode := BuildLinkedList(numSlice)
	ReadLinkedSlice(&lnode)
	// Flip1th(&lnode)
	Flip1th2(&lnode)
	ReadLinkedSlice(&lnode)
}

// Flip1th Flip adjacent linkedList 
func Flip1th(ln *LNode){
	if ln.Next == nil{
		return
	}
	cur := &LNode{}
	next := &LNode{}
	for cur, next = ln.Next, ln.Next.Next; next != nil; cur, next = ln.Next, ln.Next.Next{
		cur.Next = next.Next
		next.Next = cur
		ln.Next = next
		ln = ln.Next.Next
		if ln.Next == nil || ln.Next.Next == nil{
			break
		}
	}
}


// Flip1th2 adjacent linkedList
func Flip1th2(ln *LNode){
	if ln == nil || ln.Next == nil{
		return
	}
	cur := ln.Next
	pre := ln
	var next *LNode
	for cur != nil && cur.Next != nil{
		next = cur.Next.Next
		pre.Next = cur.Next
		cur.Next.Next = cur
		cur.Next = next
		pre = cur
		cur = next
	}
}