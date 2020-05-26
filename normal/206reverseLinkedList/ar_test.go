package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//反转一个单链表。
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

type ListNode struct {
	Val int
	Next *ListNode
}


func TestHello(t *testing.T) {
	head := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:6,
			Next:&ListNode{
				Val:9,
				Next:&ListNode{
					Val:0,
					Next:&ListNode{
						Val:9,
						Next:&ListNode{
							Val:1,
							Next:nil,
						},
					},
				},
			},
		},
	}
	t.Log(reverseList(head))
}

func reverseList(head *ListNode) *ListNode {
	result := &ListNode{}
	for head != nil {
		tmp := head.Next
		head.Next = result
		result = head
		head = tmp
	}

	return result
}
