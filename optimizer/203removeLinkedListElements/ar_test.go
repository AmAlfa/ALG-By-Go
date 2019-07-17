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

//时间最优
//国外1
func removeElements1(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	for head != nil && head.Val == val {
		head = head.Next
	}
	t := head
	for head != nil {
		if head.Next != nil && head.Next.Val == val {
			for head.Next != nil && head.Next.Val == val {
				head.Next = head.Next.Next
			}
		}
		head = head.Next
	}
	return t
}
//国外2
func removeElements2(head *ListNode, val int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	p := dummy

	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	return dummy.Next
}
//国内3
func removeElements3(head *ListNode, val int) *ListNode {
	pre:=ListNode{0,head}
	tmp:=&pre
	for tmp.Next!=nil{
		if tmp.Next.Val==val{
			tmp.Next=tmp.Next.Next
		}else{
			tmp=tmp.Next
		}
	}
	return pre.Next
}
//国内4
func removeElements4(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	var p, q *ListNode
	p = head
	q = p

	for p != nil {
		if p.Val == val {
			if p == head {
				head = p.Next
			} else {
				q.Next = p.Next
			}
			p = p.Next
			continue
		}
		q = p
		p = p.Next
	}

	return head
}
//内存最优
//国外5
func removeElements5(head *ListNode, val int) *ListNode {
	ret := &ListNode{Val: 0, Next: head}
	index := ret
	for index.Next != nil {
		if index.Next.Val == val {
			next := index.Next
			index.Next = next.Next
		} else {
			index = index.Next
		}
	}
	return ret.Next
}

//国外6
//国内7
//国内8
