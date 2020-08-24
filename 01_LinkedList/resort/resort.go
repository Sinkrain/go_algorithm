package resort

import (
	"fmt"
	. "linkedList/deduplicate"
	. "linkedList/reverse"
)


// ResortLinkedList Resort LinkedList
func ResortLinkedList(n int){
	CenterMsg("Resort LinkedList")
	numSlice := BuildNumSlice(n, 100)
	fmt.Println("NumSlice: ", numSlice)
	ln := BuildLinkedList(numSlice)
	ReadLinkedSlice(&ln)
	Resort(&ln)
	ReadLinkedSlice(&ln)

}

// Resort linkedList
func Resort(ln *LNode){
	if ln.Next == nil || ln.Next.Next == nil{
		Reverse(ln)
		return 
	}
	
	slower := ln
	faster := ln
	for faster!= nil && faster.Next != nil && faster.Next.Next != nil{
		slower = slower.Next
		faster = faster.Next.Next
	}
	tmpNode := &LNode{
		Next: slower.Next,
	}
	Reverse(tmpNode)
	slower.Next = nil
	// remove head node
	tmpNode = tmpNode.Next
	for ln != nil && tmpNode != nil{
		tmp := tmpNode.Next
		tmpNode.Next = ln.Next
		ln.Next = tmpNode
		ln = ln.Next.Next
		tmpNode = tmp
	}
}