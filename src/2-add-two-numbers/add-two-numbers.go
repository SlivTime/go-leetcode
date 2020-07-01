package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func listToInt(l *ListNode) int {
	result := 0
	multiplier := 1
	current := l

	for {
		if current == nil {
			break
		}
		result += current.Val * multiplier
		multiplier *= 10
		current = current.Next
	}
	return result
}

func intToList(n int) *ListNode {
	start := ListNode{}
	current := &start
	for n > 0 {
		nodeVal := n % 10
		current.Val = nodeVal
		n = n / 10
		if n > 0 {
			next := ListNode{}
			current.Next = &next
			current = &next
		}
	}
	return &start
}

// Simpler version with just adding two ints with some type conversions
// It is not universal due to int overflow, but it is more easy to read
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return intToList(listToInt(l1) + listToInt(l2))
}

// Common version with iteration over two linked lists
// Only this one passes the leetcode tests
func addTwoNumbersLoop(l1 *ListNode, l2 *ListNode) *ListNode {
	carryBit := 0
	result := ListNode{}
	current := &result
	currentL1 := l1
	currentL2 := l2
	for {
		if currentL1 == nil && currentL2 == nil {
			// Terminate only if both lists are empty
			if carryBit == 1 {
				current.Next = &ListNode{carryBit, nil}
			}
			break
		}
		if currentL1 == nil {
			// Create virtual zero-value Node
			currentL1 = &ListNode{}
		}
		if currentL2 == nil {
			// Create virtual zero-value Node
			currentL2 = &ListNode{}
		}
		sum := currentL1.Val + currentL2.Val + carryBit
		carryBit = 0
		if sum > 9 {
			carryBit = 1
		}
		current.Val = sum % 10

		currentL1 = currentL1.Next
		currentL2 = currentL2.Next

		if currentL1 != nil || currentL2 != nil {
			next := ListNode{}
			current.Next = &next
			current = &next
		}
	}

	return &result
}
