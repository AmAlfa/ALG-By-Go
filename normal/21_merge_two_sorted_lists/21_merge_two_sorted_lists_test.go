package merge_two_sorted_lists_test

import (
	"fmt"
	"testing"
)
// 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//
// 示例：
//
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4

type ListNode struct {
	Val int
	Next *ListNode
}

func Test(t *testing.T) {
	l1 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:4,
				Next:nil,
			},
		},
	}
	l2 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:3,
			Next:&ListNode{
				Val:4,
				Next:&ListNode{
					Val:5,
					Next:nil,
				},
			},
		},
	}
	res := mergeTwoLists(l1, l2)
	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
	// fmt.Println(res)
	// t.Log(res)

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	head := l3
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			l3.Next = l1
			l3 = l1
			l1 = l1.Next
		} else {
			l3.Next = l2
			l3 = l2
			l2 = l2.Next
		}
	}

	if l1 == nil {
		l3.Next = l2
	} else {
		l3.Next = l1
	}

	return head.Next
}
