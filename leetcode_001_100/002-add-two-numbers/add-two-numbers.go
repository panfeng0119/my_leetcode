package problem0002

// ListNode defines for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// result
	res := &ListNode{}
	// 指向当前位置的指针
	cur := res
	// 进位标记
	tmp := 0

	for l1 != nil || l2 != nil || tmp > 0 {
		sum := tmp

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		tmp = sum / 10 // 进位

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return res.Next
}
