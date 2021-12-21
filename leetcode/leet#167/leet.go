package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head==nil{
		return false
	}
	node := head
    nodeMap := make(map[*ListNode]int)
	for node.Next != nil{
		nodeMap[node]++
		if nodeMap[node]>1{
			fmt.Println("true")
			return true
		}
		node=node.Next
	}
	fmt.Println("false")
	return false
}

func main(){
	test0 := ListNode{Val: 1}
	test1 := ListNode{Val: 1}
	test2 := ListNode{Val: 1}

	test0.Next = &test1
	test1.Next = &test2
	test2.Next = nil

	hasCycle(&test0)

}
