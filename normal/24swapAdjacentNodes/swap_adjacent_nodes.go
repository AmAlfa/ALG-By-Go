package swap_adjacent_nodes

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmpNode := head
	head = head.Next
	tmpNode.Next = head.Next
	head.Next = tmpNode
	tmpNode.Next = swapPairs(tmpNode.Next)
	return head

}
//时间复杂度：O(log(n))
//空间复杂度：O(1)
