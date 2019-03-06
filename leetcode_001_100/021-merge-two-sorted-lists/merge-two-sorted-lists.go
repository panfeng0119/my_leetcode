package problem0021

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 存在空链表，直接返回另一个
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 初始化
	l := &ListNode{}
	tmp := l
	for l1 != nil && l2 != nil {
		// 谁小要谁
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	// 把剩下的补进来
	if l1 != nil {
		tmp.Next = l1
	} else {
		tmp.Next = l2
	}
	// 牛逼的一行，替换上面的功能
	// tmp.Next = map[bool]*ListNode{true: l1, false: l2}[l1 != nil]
	return l.Next
}
