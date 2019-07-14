package _validParentheses

import (
	"testing"
)
//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//
//示例：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4

type ListNode struct {
	Val int
	Next *ListNode
}

func TestHello(t *testing.T) {
	l1 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:6,
				Next:nil,
			},
		},
	}
	l2 := &ListNode{
		Val:2,
		Next:&ListNode{
			Val:3,
			Next:&ListNode{
				Val:5,
				Next:nil,
			},
		},
	}

	l3 := &ListNode{}
	currentNode := l3
	if l1.Val > l2.Val {
		for l2 != nil {
			currentNode.Next = &ListNode{
				Val:l1.Val,
			}
			currentNode = currentNode.Next
			l2 = l2.Next
		}
	} else {
		for l1 != nil {
			currentNode.Next = &ListNode{
				Val:l1.Val,
			}
			currentNode = currentNode.Next
			l1 = l1.Next
		}
	}
	currentNode.Next = nil
	otherCurrentNode := l3
	for l2 != nil {
		for otherCurrentNode != nil {
			t.Log(otherCurrentNode)
			//if otherCurrentNode.Next != nil {
			//	if otherCurrentNode.Next != nil && l2.Val < otherCurrentNode.Next.Val {
			//		newNode := &ListNode{
			//			Val:l2.Val,
			//		}
			//		newNode.Next = otherCurrentNode.Next
			//		otherCurrentNode = otherCurrentNode.Next
			//	}
			//}
			if otherCurrentNode.Next == nil {
				otherCurrentNode.Next = &ListNode{
					Val:l2.Val,
				}
				otherCurrentNode = otherCurrentNode.Next
			}
			otherCurrentNode = otherCurrentNode.Next
		}

		l2 = l2.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//
//}
