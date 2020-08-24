package loop

import (
	"fmt"
	"math/rand"
	. "linkedList/deduplicate"
)


// CheckIsLoop Checked the linkedList whether it`s looped 
func CheckIsLoopLinkedList(n int){
	CenterMsg("CheckIsLoop for LinkedList")
	numSlice := BuildNumSlice(n, 100)
	fmt.Println("NumSlice: ", numSlice)
	lNode := BuildLinkedList(numSlice)
	ReadLinkedSlice(&lNode)
	BuildLoopLinkedList(&lNode, n)
	readLoopLinkedList(&lNode, 30)
	CheckIsLoop(&lNode)

}

// CheckIsLoop check whether is loop
func CheckIsLoop(ln *LNode){
	if ln.Next == nil || ln.Next.Next == nil{
		return
	}

	slow, faster := ln.Next, ln.Next.Next
	for slow.Next != nil && faster.Next.Next != nil{

		if slow == faster {
			fmt.Println("\nIs Looped")
			ReadLooped(slow, faster)
			return
		} else if faster == nil{
			fmt.Println("\nNot`s Looped")
			return
		}
		slow = slow.Next
		faster = faster.Next.Next 
	}
	
}

// ReadLooped read linkedList
func ReadLooped(slow *LNode, faster *LNode){
	fmt.Println("LoopedLinked: ")
	slow = slow.Next
	for slow != faster{
		fmt.Printf("-> %d", slow.Num)
		slow = slow.Next
	}
	fmt.Printf("-> %d\n", slow.Num)
}

// BuildLoopLinkedList Build Loop LinkedList
func BuildLoopLinkedList(ln *LNode, n int){
	// range number
	k := rand.Intn(n)
	for k < 1 && n > 1 {
		k = rand.Intn(n)
	}
	pointer := &LNode{}
	for ln.Next != nil{
		if k == 0 {
			pointer.Next = ln.Next
		}
		ln = ln.Next
		k --
	}
	ln.Next = pointer.Next
}


// readLoopLinkedList print LoopLinkedList
func readLoopLinkedList(ln *LNode, n int){
	fmt.Println("LoopLinkedList:")
	for ; n > 0; n--{
		fmt.Printf("-> %d ", ln.Num)
		ln = ln.Next
	}
}