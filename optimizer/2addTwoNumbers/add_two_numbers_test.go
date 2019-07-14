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
	t.Log(addTwoNumbers1(l1, l2).Val)
	t.Log(addTwoNumbers1(l1, l2).Next.Val)
	t.Log(addTwoNumbers1(l1, l2).Next.Next.Val)
}
//最快的两个
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	head := &ListNode{0, nil}
	first := head
	suc := 0
	sum := 0
	for l1 != nil && l2 != nil {
		node := &ListNode{0, nil}
		sum = l1.Val + l2.Val + suc
		node.Val = sum % 10
		suc = sum / 10
		first.Next = node
		first = node
		l1 = l1.Next //不要忘记这一步
		l2 = l2.Next
	}

	for l1 != nil {
		node := &ListNode{0, nil}
		sum = l1.Val + suc
		node.Val = sum % 10
		suc = sum / 10
		first.Next = node
		first = node
		l1 = l1.Next //不要忘记这一步
	}

	for l2 != nil {
		node := &ListNode{0, nil}
		sum = l2.Val + suc
		node.Val = sum % 10
		suc = sum / 10
		first.Next = node
		first = node
		l2 = l2.Next //不要忘记这一步
	}

	if suc != 0 {
		node := &ListNode{suc, nil}
		first.Next = node
		first = node
	}
	return head.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var newRoot *ListNode
	keep := true
	aux1 := l1
	aux2 := l2
	increment := 0

	for keep {
		if aux1 != nil && aux2 != nil {
			node := sumTwoNodes(&aux1, &aux2, &increment)
			node.Next = newRoot
			newRoot = node
		} else if aux1 != nil {
			node := sumTwoNodes(&aux1, nil, &increment)
			node.Next = newRoot
			newRoot = node
		} else if aux2 != nil {
			node := sumTwoNodes(nil, &aux2, &increment)
			node.Next = newRoot
			newRoot = node
		} else if increment > 0 {
			node := new(ListNode)
			node.Val = increment
			node.Next = newRoot
			newRoot = node
			increment = 0
		} else {
			keep = false
		}
	}

	return reverse(newRoot)
}

func sumTwoNodes(l1 **ListNode, l2 **ListNode, increment *int) *ListNode {
	node := new(ListNode)
	node.Val = *increment

	if l1 != nil && *l1 != nil {
		node.Val += (*l1).Val
	}

	if l2 != nil && *l2 != nil {
		node.Val += (*l2).Val
	}

	*increment = 0
	if node.Val/10 >= 1 {
		node.Val %= 10
		*increment++
	}

	if l1 != nil && *l1 != nil {
		*l1 = (*l1).Next
	}

	if l2 != nil && *l2 != nil {
		*l2 = (*l2).Next
	}

	return node
}

func reverse(l *ListNode) *ListNode {
	var newRoot *ListNode
	aux := l

	for aux != nil {
		node := new(ListNode)
		node.Val = aux.Val
		node.Next = newRoot
		newRoot = node
		aux = aux.Next
	}

	return newRoot
}


//内存最少的两个
func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	node1 := l1
	node2 := l2

	var prevNode1 *ListNode
	carry := 0

	for node1 != nil && node2 != nil {
		val := node1.Val + node2.Val + carry
		carry = 0
		if val >= 10 {
			carry = 1
			val -= 10
		}

		node1.Val = val

		prevNode1 = node1
		node1 = node1.Next
		node2 = node2.Next
	}

	if node1 == nil && node2 != nil {
		prevNode1.Next = node2
		node1 = node2
	}

	for node1 != nil {
		val := node1.Val + carry
		carry = 0
		if val >= 10 {
			carry = 1
			val -= 10
		}
		node1.Val = val

		prevNode1 = node1
		node1 = node1.Next
	}

	if carry > 0 {
		prevNode1.Next = &ListNode{
			Val: carry,
		}
	}

	return l1
}


func addTwoNumbers4(l1 *ListNode, l2 *ListNode) *ListNode {
	res2 := l1
	for {
		if l2 == nil {
			break
		}
		if l1 == nil {
			l1 = l2
			break
		}
		l1.Val += l2.Val
		l2 = l2.Next
		if l1.Next == nil {
			l1.Next = l2
			break
		}
		l1 = l1.Next
	}
	l1 = res2
	for {
		if l1 == nil {
			break
		}
		if l1.Val >= 10 {
			if l1.Next == nil {
				l1.Next = &ListNode{}
			}
			l1.Next.Val += l1.Val / 10
			l1.Val = l1.Val % 10
		}
		l1 = l1.Next
	}
	return res2

	i1 := val(l1)
	i2 := val(l2)
	var res []int
	if len(i1) > len(i2) {
		res = add(i1, i2)
	} else {
		res = add(i2, i1)
	}
	return fromInt(res)
}

func fromInt(i []int) *ListNode {
	res := &ListNode{}
	node := res
	var prev *ListNode
	for _, v := range i {
		node.Val = v
		node.Next = &ListNode{}
		prev = node
		node = node.Next
	}
	if prev != nil {
		prev.Next = nil
	}

	return res
}

func val(l *ListNode) []int {
	result := make([]int, 0, 50)
	for {
		if l == nil {
			break
		}

		result = append(result, l.Val)
		l = l.Next
	}

	return result
}

// i1 should be longer or equal to i2
func add(i1, i2 []int) []int {
	result := i1
	for i, v := range i2 {
		result[i] += v
	}
	result = append(result, 0)
	for i, v := range result {
		if v >= 10 {
			result[i] = v % 10
			result[i+1] += v/10
		}
	}
	if result[len(result) -1] == 0 {
		result = result[0:len(result)-1]
	}
	return result
}
