package deduplicate


// DeduplicateSlice2 deduplicate
func DeduplicateSlice(ls *LNode){
	if ls == nil || ls.Next == nil{
		return
	}
	outerCur := ls.Next
	var innerCur *LNode
	var innerPre *LNode
	for ; outerCur != nil; outerCur=outerCur.Next{
		for innerCur, innerPre = outerCur.Next, outerCur; innerCur!=nil;{
			if outerCur.Num == innerCur.Num{
				innerPre.Next = innerCur.Next
			} else {
				innerPre = innerCur
			}
			innerCur = innerCur.Next
		}
	}
}

// DeduplicateSlice deduplicate slice 
func DeduplicateSlice2(ls *LNode){
	for {
		cNum := ls.Num
		if ls.Next == nil{
			return
		}
		
		pre := ls
		var cur *LNode
		for {
			if pre.Next == nil{
				break
			}
			cur = pre.Next

			if cur.Num == cNum{
				if cur.Next == nil{
					break
				}
				pre.Next = cur.Next
				continue				
			}
			pre = pre.Next
		}
		ls = ls.Next
	}

}
