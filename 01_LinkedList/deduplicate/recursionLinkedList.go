package deduplicate


// RecursionDeduplicate Recursion deduplication
func RecursionDeduplicate(rl *LNode)*LNode{
	if rl == nil || rl.Next == nil{
		return rl
	}
	var pointer *LNode
	cur := rl
	// 对 rl.Next 后面的链表去重
	rl.Next = RecursionDeduplicate(rl.Next)
	pointer = rl.Next

	for pointer != nil{
		if rl.Num == pointer.Num{
			cur.Next = pointer.Next
			pointer = cur.Next
		} else {
			pointer = pointer.Next
			cur = cur.Next
		}
	}
	return rl

}
