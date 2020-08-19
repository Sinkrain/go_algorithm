package deduplicate

import (
	"fmt"
	"math/rand"
)

type LNode struct{
	Num int
	Next * LNode
}

// ReadLinkedSlice read linkedSlice
func ReadLinkedSlice(l *LNode){
	for {
		if l.Next == nil{
			fmt.Println("")
			return
		}
		l = l.Next
		fmt.Printf(" -> %d", l.Num)
	}
}

// BuildNumSlice create numSlice
func BuildNumSlice(n int, limit int) (numSlice []int){
	for i:=0; i<n; i++{
		numSlice = append(numSlice, rand.Intn(limit))
	}
	return numSlice
}

// BuildLinkedList Create linkedList by slice
func BuildLinkedList(numSlice []int)LNode{
	head := LNode{}
	for _,n := range numSlice{
		newLNode := &LNode{
			Num: n,
		}
		newLNode.Next = head.Next
		head.Next = newLNode
	}
	return head
}
