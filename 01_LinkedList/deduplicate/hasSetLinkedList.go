package deduplicate


// HasSetLinkedList dedulicate linkedList
func HasSetLinkedList(hl *LNode){
	tmpMap := map[int]bool{}
	for hl != nil && hl.Next != nil{
		cur := hl.Next
		if _, e := tmpMap[cur.Num]; !e{
			tmpMap[cur.Num] = true
			continue
		}
		hl.Next = cur.Next
		hl = hl.Next
	}


}