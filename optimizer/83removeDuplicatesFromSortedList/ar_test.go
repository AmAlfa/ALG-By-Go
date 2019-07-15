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
//时间最优
//国外1
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	list := &ListNode{Next: head}

	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}

	return list.Next
}
//国外2
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dict := make(map[int]bool)
	prev := head
	cur := head.Next
	dict[head.Val] = true
	for cur != nil {
		if _, ok := dict[cur.Val]; !ok {
			dict[cur.Val] = true
			prev = prev.Next
		} else {
			prev.Next = cur.Next
		}
		cur = cur.Next
	}
	return head
}
//国内3
func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	for fast:=head.Next; fast != nil; fast=fast.Next {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		} else if fast.Next == nil {
			slow.Next = nil
		}
	}
	return head
}
//国内4
func deleteDuplicates4(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for {
		if cur == nil ||cur.Next == nil {
			break
		}
		for {
			if cur.Next == nil||  cur.Next.Val != cur.Val  {
				break
			}
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}
//内存最优
//国外5
func deleteDuplicates5(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}

	out := head
	last := out
	head = head.Next
	out.Next = nil
	lastElem := out.Val

	for head != nil {
		if lastElem != head.Val {
			lastElem = head.Val
			last.Next = head
			head = head.Next
			last = last.Next
			last.Next = nil
		} else {
			head = head.Next
		}
	}

	return out
}
//国外6
func deleteDuplicates6(head *ListNode) *ListNode {
	begin := head
	for head != nil && head.Next != nil {
		for head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return begin
}
//国内7
//国内8
