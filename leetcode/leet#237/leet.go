package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


func printLinked(head *ListNode){
	node := head
	for node!=nil{
		fmt.Println(node.Val)
		node=node.Next
	}
}

func deleteNode(node *ListNode) {
	for node!=nil{
		if node.Next.Next == nil{
			node.Val= node.Next.Val
			node.Next = nil
			return
		}
		node.Val = node.Next.Val
		node = node.Next
	}
    
}

func main(){
	head := ListNode{Val: 0, Next: nil}
	node1 := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 4, Next: nil}

	head.Next = &node1
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4

	deleteNode(&node1)
	printLinked(&head)
}
