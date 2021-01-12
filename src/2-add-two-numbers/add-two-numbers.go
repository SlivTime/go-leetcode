package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry, value int
	result := &ListNode{}
	current := result

	for l1 != nil || l2 != nil {
		value = carry
		if l1 != nil {
			value += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			value += l2.Val
			l2 = l2.Next
		}
		carry = value / 10
		value = value % 10

		current.Next = &ListNode{value, nil}

		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{carry, nil}
	}
	return result.Next
}
