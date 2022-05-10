package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	left := l1
	right := l2
	total := res
	mind := 0
	mock := &ListNode{}

	for {
		v := left.Val + right.Val + mind
		total.Val = v % 10
		if v >= 10 {
			mind = 1
		} else {
			mind = 0
		}

		left = left.Next
		right = right.Next
		if left == nil && right == nil {
			break
		}

		if left == nil {
			left = mock
		}

		if right == nil {
			right = mock
		}

		total.Next = &ListNode{}
		total = total.Next
	}

	if mind >= 1 {
		total.Next = &ListNode{}
		total.Next.Val = 1
	}

	return res
}
