package flip


import (
	// "fmt"
	. "linkedList/deduplicate"
	. "linkedList/reverse"
)

// FlipNthGroupLinkedList Filp LinkedList by Nth Group
func FlipNthGroupLinkedList(n int){
	CenterMsg("Flip Nth Group LinkedList")
	numSlice := BuildNumSlice(n, 100)
	lnode := BuildLinkedList(numSlice)
	ReadLinkedSlice(&lnode)
	FlipNth(&lnode, 3)
	ReadLinkedSlice(&lnode)
}

// FlipNth Flip linkedList by nth group
func FlipNth(ln *LNode, n int){
	pre := ln
	cur_l := ln.Next
	cur_r := &LNode{}
	next := ln

	for pre.Next != nil && next != nil{
		cur := pre.Next
		for i:= 0; i <n; i++{
			if cur.Next == nil{
				return
			}
			cur = cur.Next
		}
		cur_r = cur.Next
		cur.Next = nil
		Reverse(cur)
		cur_l.Next = cur_r
		pre = cur_l
		cur_l = cur_r
	}
}

