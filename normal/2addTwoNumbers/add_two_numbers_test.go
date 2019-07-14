package two_sum

import (
	"testing"
)

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

type ListNode struct {
	Val int
	Next *ListNode
}

func TestHello(t *testing.T) {
	l1 := &ListNode{
		Val:5,
		Next:&ListNode{
			Val:8,
			Next:&ListNode{
				Val:3,
			},
		},
	}

	l2 := &ListNode{
		Val:5,
		Next:&ListNode{
			Val:6,
			Next:&ListNode{
				Val:4,
			},
		},
	}
	t.Log(addTwoNumbers(l1, l2).Val)
	t.Log(addTwoNumbers(l1, l2).Next.Val)
	t.Log(addTwoNumbers(l1, l2).Next.Next.Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	currentNode := l3
	carry := 0
	for l1 != nil || l2 != nil {
		sumNum := 0
		if l1 == nil {
			sumNum += l2.Val
		} else if l2 == nil {
			sumNum += l1.Val
		} else {
			sumNum = l1.Val + l2.Val
		}

		sumNum += carry
		carry = sumNum / 10
		currentNode.Next = &ListNode{
			Val: sumNum % 10,
		}
		currentNode = currentNode.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		currentNode.Next = &ListNode{
			Val:carry,
			Next:nil,
		}
	}

	return l3.Next
}
