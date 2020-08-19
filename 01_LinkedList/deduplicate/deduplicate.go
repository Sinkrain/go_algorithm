package deduplicate

import "fmt"


// DeduplicateLinkedList deduplicate linkedList
func DeduplicateLinkedList(n int){
	numSlice := BuildNumSlice(n, 100)
	fmt.Println("New Slice: ", numSlice)
	baseLinkSlice := BuildLinkedList(numSlice)
	numLinkedSlice := BuildLinkedList(numSlice)
	numLinkedSlice2 := BuildLinkedList(numSlice)
	ReadLinkedSlice(&numLinkedSlice)
	
	// dedulication linkedList by lterate
	fmt.Println("Letrate Deduplicate Slice:")
	DeduplicateSlice(&numLinkedSlice)
	ReadLinkedSlice(&numLinkedSlice)
	
	fmt.Println("Recursion Deduplicate Slice")
	RecursionDeduplicate(&baseLinkSlice)
	ReadLinkedSlice(&baseLinkSlice)
	
	fmt.Println("Hasset Deduplicate Slice")
	HasSetLinkedList(&numLinkedSlice2)
	ReadLinkedSlice(&baseLinkSlice)

}


