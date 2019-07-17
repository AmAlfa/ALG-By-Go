package _removeDuplicatesFromSortedArray

import (
	"testing"
)
//删除链表中等于给定值 val 的所有节点。
//
//示例:
//
//输入: 1->2->6->3->4->5->6, val = 6
//输出: 1->2->3->4->5
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
	val := 2

	t.Log(removeElements(head, val))
}

func removeElements(head *ListNode, val int) *ListNode {
	result := &ListNode{}
	if head == nil {
		return head
	}
	resultHead := result
	for head != nil {
		if head.Val != val {
			resultHead.Next = &ListNode{
				Val:head.Val,
			}
			resultHead = resultHead.Next
		}
		head = head.Next
	}
	resultHead.Next = nil
	return result.Next
}
