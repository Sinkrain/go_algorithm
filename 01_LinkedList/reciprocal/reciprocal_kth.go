package reciprocal

import (
	"fmt"
	. "linkedList/deduplicate"
)

// ReciprovalLinkedList Reciproval LinkedList
func ReciprovalLinkedList(n int, k int){
	CenterMsg(fmt.Sprintf("Reciproval %dth LinkedList", k))
	if k > n {
		fmt.Println("Please Input right num of n and k which k <= n")
		return
	}

	numSlice := BuildNumSlice(n, 100)
	fmt.Println("Num Slice: ", numSlice)
	lNode := BuildLinkedList(numSlice)
	ReadLinkedSlice(&lNode)
	kth := ReciprovalNth(&lNode, k)
	fmt.Printf("The reciproval %dth: %d\n", k, kth)
}

// ReciprovalNth get the reciproval Nth num of linkedList
func ReciprovalNth(lnode *LNode, k int)int{
	faster := lnode
	for ; k > 0; k--{
		if faster == nil{
			panic("LinkedList is too short")
		}
		faster = faster.Next
	}
	for faster != nil{
		lnode = lnode.Next
		faster = faster.Next
	}
	return lnode.Num
}