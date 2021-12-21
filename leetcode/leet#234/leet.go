package main

//import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	nodes := make(map[int]int)
	index := 0
    for head != nil{
		nodes[index] = head.Val
		//fmt.Printf("index: %d, ",index)
		//fmt.Println(nodes[index])
		index++
		head = head.Next
	}
	if index%2==0{
		//fmt.Println("%2")
		for i:=0; i<=index/2; i++{
			if nodes[i] != nodes[index-1]{
				//fmt.Println("false")
				return false
			}
			index--
		}
	}else{
		//fmt.Println("!%2")
		for i:=0; i<=index; i++{
			//fmt.Printf("comparing %d, %d to %d, %d\n",i, nodes[i], index, nodes[index-1])
			if nodes[i] != nodes[index-1]{
				//fmt.Println("false")
				return false
			}
			index--
		}
	}
	
	//fmt.Println("True")
	return true
}

func main(){
	head := ListNode{Val: 1, Next: nil}
	node1 := ListNode{Val: 2, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	//node3 := ListNode{Val: 1, Next: nil}

	head.Next = &node1
	node1.Next = &node2
	//node2.Next = &node3

	isPalindrome(&head)
}
