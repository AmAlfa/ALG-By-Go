package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//
//示例 1:
//
//输入: 1->1->2
//输出: 1->2
//示例 2:
//
//输入: 1->1->2->3->3
//输出: 1->2->3

type ListNode struct {
	Val  int
	Next *ListNode
}
func TestHello(t *testing.T) {
	head := &ListNode{
		Val:0,
		Next:&ListNode{
			Val:0,
			Next:&ListNode{
				Val:0,
				Next:nil,
			},
		},
	}

	t.Log(deleteDuplicates(head).Val)
}

func deleteDuplicates(head *ListNode) *ListNode {
	l3 := &ListNode{}
	current := l3
	isFirst := 1
	for head != nil {
		if isFirst == 1 {
			current.Next = &ListNode{
				Val:head.Val,
			}
			current = current.Next
			isFirst = 0
		}
		if head.Val != current.Val {
			current.Next = &ListNode{
				Val:head.Val,
			}
			current = current.Next
		}
		head = head.Next
	}
	return l3.Next
}
