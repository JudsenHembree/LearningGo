package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil{
		return head
	}

	nodeHead := new(ListNode)
	newHead := true
	node := head
	newNode := head

	for node.Next.Next != nil{
		node = node.Next
	}

	for node != head{
		if newHead{
			nodeHead = node.Next
			node.Next = nil
			newHead = false
			newNode = nodeHead
			//fmt.Println("here")
		}else{
			newNode.Next = node.Next
			node.Next = nil
			newNode = newNode.Next
		}
		node = head
		for node.Next.Next != nil{
			node = node.Next
		}
	}
	if newHead{
		nodeHead = head.Next
		nodeHead.Next = head
		nodeHead.Next.Next = nil


	}else{
		newNode.Next = head.Next
		newNode.Next.Next = head
		newNode.Next.Next.Next = nil
	}

	//printLinked(nodeHead)
	
	return nodeHead
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
	//node1 := ListNode{Val: 2, Next: nil}
	//node2 := ListNode{Val: 2, Next: nil}
	//node3 := ListNode{Val: 1, Next: nil}
	//node4 := ListNode{Val: 1, Next: nil}

	//head.Next = &node1
	//node1.Next = &node2
	//node2.Next = &node3
	//node3.Next = &node4

	reverseList(&head)

}
