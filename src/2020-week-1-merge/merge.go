package merge

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result, lastNode *ListNode

	appendAndShift := func(value *ListNode) {
		if result == nil {
			result = value
			lastNode = value
		} else {
			lastNode.Next = value
			lastNode = lastNode.Next
		}

	}

	for l1 != nil || l2 != nil {
		if l1 == nil {
			// append l2
			appendAndShift(l2)
			l2 = l2.Next
		} else if l2 == nil {
			// append l1
			appendAndShift(l1)
			l1 = l1.Next
		} else {
			// compare l1 and l2 and append smallest
			if l1.Val <= l2.Val {
				appendAndShift(l1)
				l1 = l1.Next
			} else {
				appendAndShift(l2)
				l2 = l2.Next
			}
		}
	}
	return result
}
