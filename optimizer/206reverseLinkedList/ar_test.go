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
//时间最优
//国外1
//国外2
//国内3
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	next := head.Next
	for next != nil {
		nextNext := next.Next
		next.Next = p
		p = next
		next = nextNext
		//next, next.Next, p = next.Next, p, next
	}

	head.Next = nil
	return p
}
//国内4
//内存最优
//国外5
func reverseList5(head *ListNode) *ListNode {
	var prev *ListNode = nil
	cur := head
	var next *ListNode = nil

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}
//国外6
func reverseList7(head *ListNode) *ListNode {

	p := head
	var res *ListNode
	for p!= nil{
		t := p.Next
		p.Next = res
		res = p
		p = t
	}

	return res
}

//国内7
//国内8
