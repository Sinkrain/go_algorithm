package sum


import (
	"fmt"
	. "linkedList/deduplicate"
)

// SumLinkedListByCount add linkedList by count
func SumLinkedListByCount(l LNode, r LNode){
	lNum := CountSinglyLinkedList(&l)
	rNum := CountSinglyLinkedList(&r)
	fmt.Printf("\nLinkedNum(%d + %d)=: %d", lNum, rNum, lNum + rNum)

}

// CountSinglyLinkedList calculate the sum of singly LinkedList 
func CountSinglyLinkedList(l *LNode)(num int){
	std :=1
	cur := l.Next
	for cur != nil{
		num += cur.Num*std
		std *= 10
		if cur.Next == nil{
			break
		}
		cur = cur.Next
	}
	return num
}