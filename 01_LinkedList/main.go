package main

import (
	"linkedList/novel"
	"linkedList/reverse"
	"linkedList/deduplicate"
	"linkedList/sum"
)


func main(){

	// Noval example for LinkedList
	newNovel := novel.WriteNovel()
	novel.ReadNovel(newNovel)

	// ReverseLinkedList
	reverse.ReverseLinkedList(6)

	// DeduplicationLinkedList 
	deduplicate.DeduplicateLinkedList(32)

	// Sum linkedList
	sum.SumLinkedList(3)
}