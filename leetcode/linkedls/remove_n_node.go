package main

import "fmt"

func main() {
	node5 := ListNode{5, nil}
	node4 := ListNode{4, &node5}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}
	res := removeNthFromEnd(&node1, 5)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	var prv *ListNode
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast != nil {
		fast = fast.Next
		prv = slow
		slow = slow.Next
	}
	if slow == head {
		return head.Next
	} else {
		prv.Next = slow.Next
		return head
	}
}
