package sum


import (
	"fmt"
	. "linkedList/deduplicate"
	. "linkedList/reverse"
)


// SumLinkedListByAdd sum LinkedList by add
func SumLinkedListByAdd(l LNode, r LNode){
	fmt.Println("\nSum LinkedList by Add")
	numLinkedList := AddLinkedList(&l, &r)
	ReadLinkedSlice(numLinkedList)
	Reverse(numLinkedList)
	num := CountSinglyLinkedList(numLinkedList)
	fmt.Println("LinkedListSum: ", num)

}

// AddLinkedList add linkedList
func AddLinkedList(l *LNode, r *LNode)*LNode{
	s := &LNode{}
	lc := l.Next
	rc := r.Next
	ec := 0

	for lc != nil && rc != nil {
		cNum := lc.Num + rc.Num
		tmp := &LNode{
			Num: cNum + ec,
		}
		tmp.Next = s.Next
		s.Next = tmp
	
		if tmp.Num > 9{
			ec = tmp.Num / 10
			tmp.Num= tmp.Num % 10
		} else {
			ec = 0
		}

		lc = lc.Next
		rc = rc.Next
	}

	if lc == nil && rc == nil && ec!=0{
		s.Next = &LNode{
			Num: 1,
		}
	}

	AddLeftLinkedList(lc, s, ec)
	AddLeftLinkedList(rc, s, ec)
	return s
}

// AddLeftLinkedList add the left linked node for the lengther
func AddLeftLinkedList(l *LNode, s *LNode, ec int){
	for l != nil{
		tmp := &LNode{
			Num: l.Num + ec,
		}
		if tmp.Num > 9{
			tmp.Num -= 10
			ec = 1
		} else {
			ec = 0
		}

		tmp.Next = s.Next
		s.Next = tmp

		l = l.Next
	}
}
