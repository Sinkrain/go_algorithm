package sum


import (
	"fmt"
	. "linkedList/deduplicate"
)


// SumLinkedListByAdd sum LinkedList by add
func SumLinkedListByAdd(l LNode, r LNode){
	fmt.Println("\nSum LinkedList by Add")
	numLinkedList := AddLinkedList(&l, &r)
	num := CountSinglyLinkedList(numLinkedList)
	fmt.Println("LinkedListSum: ", num)

}

// AddLinkedList add linkedList
func AddLinkedList(l *LNode, r *LNode)*LNode{
	s := &LNode{}
	lc := l.Next
	rc := r.Next
	sl := true
	sr := true
	ec := 0
	std := 1

	for sl || sr {
		lNum := 0
		rNum := 0

		if sl {
			lNum = lc.Num
		}
		if sr {
			rNum = rc.Num
		}

		cNum := lNum + rNum
		tmp := &LNode{
			Num: cNum % 10 + ec,
		}
		tmp.Next = s.Next
		s.Next = tmp
	
		if cNum > 10{
			ec = 1
		} else {
			ec = 0
		}

		if lc.Next == nil{
			sl = false
		} else {
			lc = lc.Next
		}
		if rc.Next == nil{
			sr = false
		} else {
			rc = rc.Next
		}
		std *=10
	}
	return s
}

