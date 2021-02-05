package removeNthNodeFromEndOfList

import (
	"fmt"
	"testing"
)

// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
// 示例：
//
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
//
// 当删除了倒数第二个节点后，链表变为 1->2->3->5.
// 说明：
//
// 给定的 n 保证是有效的。
//
// 进阶：
//
// 你能尝试使用一趟扫描实现吗？

type ListNode struct {
	Val int
	Next *ListNode
}


func Test(t *testing.T) {
	head := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:6,
			Next:&ListNode{
				Val:9,
				Next:&ListNode{
					Val:2,
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
	res := removeNthFromEnd(head, 2)
	fmt.Println(res)
	t.Log(res)

}

// BF解法
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// if n == 0 {
// 		return head
// 	}
//
//	listLen := 1
//	for current := head; current.Next != nil; current = current.Next {
//		listLen++
//	}
//	if listLen == n {
//		return head.Next
//	}
//
//	befNode := &ListNode{}
//	coordinate := listLen - n
//	for current, tag := head, 1; tag <= coordinate; current, tag = current.Next, tag + 1 {
//		if tag == coordinate {
//			befNode = current
//		}
//	}
//
//	befNode.Next = befNode.Next.Next
//
//	return head

// 哨兵
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
//    if n == 0 {
//		return head
//	}
//
//	listLen := 1
//	for cur := head; cur.Next != nil; cur = cur.Next {
//		listLen++
//	}
//	if listLen <= n {
//		return head.Next
//	}
//
//	guard := &ListNode{Next: head}
//	tag := listLen - n
//	befNode := &ListNode{}
//	for cur, i := guard, 0; i <= tag; cur,i = cur.Next, i + 1 {
//		befNode = cur
//	}
//
//	befNode.Next = befNode.Next.Next
//
//	return guard.Next
// }

// 双指针
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	if n == 0 {
// 		return head
// 	}
//
// 	guard := &ListNode{Next: head}
//
// 	before := guard
// 	end := guard
// 	for i := 0; i < n && before.Next != nil; i++ {
// 		before = before.Next
// 	}
// 	for before.Next != nil {
// 		end = end.Next
// 		before = before.Next
// 	}
//
// 	end.Next = end.Next.Next
// 	return guard.Next
// }

// 哈希表
// *ListNode
func removeNthFromEnd(head *ListNode, n int) interface{} {
	if n == 0 {
		return head
	}
	listMap := make(map[int]*ListNode)
	listLen := 1
	for cur, key := head, 1; cur != nil; cur, key = cur.Next, key + 1 {
		listMap[key] = cur
		listLen = key
	}
	if listLen <= n {
		return head.Next
	}

	befKey := listLen - n
	listMap[befKey].Next = listMap[befKey].Next.Next

	return head
}
