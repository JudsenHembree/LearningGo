package main

import (
	"fmt"
)

//Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil{
		return head
	}
	for head.Val == val{
		if head.Next == nil{
			head=nil
			//fmt.Println("returning nil")
			return head
		}
		head = head.Next
	}

	tempNode := head
	for tempNode.Next != nil{
		if tempNode.Next.Val == val{
			tempNode.Next=tempNode.Next.Next
		}else{
			tempNode=tempNode.Next
		}
	}
	//printLinked(head)
	return head
}

func printLinked(head *ListNode){
	node := head
	for node!=nil{
		fmt.Println(node.Val)
		node=node.Next
	}
}

func main(){
	head := ListNode{Val: 1, Next: nil}
	node1 := ListNode{Val: 2, Next: nil}
	node2 := ListNode{Val: 1, Next: nil}
	node3 := ListNode{Val: 1, Next: nil}

	head.Next = &node1
	node1.Next = &node2
	node2.Next = &node3

	removeElements(&head, 2)
}
