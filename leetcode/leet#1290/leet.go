package main

import (
	"fmt"
	"math"
)

//Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var sol float64
	indexedMap := make(map[int]int)

	index := 0
	node := head
	for node!=nil{
		indexedMap[index] = node.Val
		//fmt.Println(index, " ", node.Val)
		node=node.Next
		index++
	}

	formap := 0
	for index>=0{
		temp := indexedMap[formap]
		//fmt.Println(temp, "* 2^", index)
		sol += (float64(temp)*math.Pow(float64(2), float64(index-1)))
		//fmt.Println(index, formap)
		index--
		formap++
	}

	//fmt.Println("final ", int(sol))
	return int(sol)
}

func main(){
	head := ListNode{Val: 1, Next: nil}
	node1 := ListNode{Val: 0, Next: nil}
	node2 := ListNode{Val: 1, Next: nil}
	node3 := ListNode{Val: 1, Next: nil}

	head.Next = &node1
	node1.Next = &node2
	node2.Next = &node3

	getDecimalValue(&head)
}
