package remove_duplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var currentValue, counter int
	var result, lastNode *ListNode

	if head == nil {
		return nil
	}

	appendAndShift := func(value *ListNode) {
		if result == nil {
			result = value
			lastNode = value
		} else {
			lastNode.Next = value
			lastNode = lastNode.Next
		}
	}

	currentValue = head.Val
	counter = 1
	head = head.Next

	for head != nil {
		if head.Val == currentValue {
			// just increase counter
			counter += 1
		} else {
			// reset counter and append or drop
			if counter == 1 {
				appendAndShift(&ListNode{currentValue, nil})
			}
			counter = 1
		}
		currentValue = head.Val
		head = head.Next
	}

	if counter == 1 {
		appendAndShift(&ListNode{currentValue, nil})
	}

	return result
}
