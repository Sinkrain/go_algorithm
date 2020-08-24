package deduplicate

import "fmt"


// DeduplicateLinkedList deduplicate linkedList
func DeduplicateLinkedList(n int){
	CenterMsg("Deduplicate LinkedList")
	numSlice := BuildNumSlice(n, 100)
	fmt.Println("New Slice: ", numSlice)
	baseLinkSlice := BuildLinkedList(numSlice)
	numLinkedSlice := BuildLinkedList(numSlice)
	numLinkedSlice2 := BuildLinkedList(numSlice)
	ReadLinkedSlice(&numLinkedSlice)
	
	// dedulication linkedList by lterate
	CenterMsg("Letrate Deduplicate Slice")
	DeduplicateSlice(&numLinkedSlice)
	ReadLinkedSlice(&numLinkedSlice)
	
	CenterMsg("Recursion Deduplicate Slice")
	RecursionDeduplicate(&baseLinkSlice)
	ReadLinkedSlice(&baseLinkSlice)
	
	CenterMsg("Hasset Deduplicate Slice")
	HasSetLinkedList(&numLinkedSlice2)
	ReadLinkedSlice(&baseLinkSlice)

}


