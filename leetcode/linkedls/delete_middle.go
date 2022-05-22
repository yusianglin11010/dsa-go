package main

import "fmt"

var node5 ListNode = ListNode{5, nil}
var node4 ListNode = ListNode{4, &node5}

var node3 ListNode = ListNode{3, &node4}
var node2 ListNode = ListNode{2, &node3}
var node1 ListNode = ListNode{1, &node2}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	res := deleteMiddle(&node1_1)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	slow, fast := head, head
	var prev *ListNode = nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}
	prev.Next = slow.Next
	return head
}
