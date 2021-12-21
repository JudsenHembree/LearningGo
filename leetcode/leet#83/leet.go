package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	Nodes := make(map[int]int)
	node := head

	if head == nil || head.Next == nil{
		return head
	}

	if head.Next.Next == nil{
		if head.Next.Val == head.Val{
			head.Next = nil
		}
		return head
	}
	Nodes[head.Val]=0

	for node.Next.Next!=nil{
		_, there := Nodes[node.Next.Val]
		if there {
			//fmt.Println("already there")
			node.Next = node.Next.Next
		}else{
			Nodes[node.Next.Val]=0
			node = node.Next
		}
	}

	if node.Next.Val == node.Val{
		node.Next = nil
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
	head := ListNode{Val: 0, Next: nil}
	node1 := ListNode{Val: 0, Next: nil}
	node2 := ListNode{Val: 0, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 3, Next: nil}

	head.Next = &node1
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4

	deleteDuplicates(&head)
}
