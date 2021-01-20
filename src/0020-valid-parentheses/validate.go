package parensvalidation

import "errors"

type RuneStack struct {
	items  []rune
	Length int
}

// Push item to stack
func (s *RuneStack) Push(item rune) {
	if len(s.items) > s.Length {
		s.items[s.Length] = item
	} else {
		s.items = append(s.items, item)
	}
	s.Length++
}

// Pop from stack
func (s *RuneStack) Pop() (rune, error) {
	if s.Length > 0 {
		s.Length--
		return s.items[s.Length], nil
	}
	return 0, errors.New("Empty")
}

func isValid(s string) bool {
	stack := new(RuneStack)
	parensMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, chr := range s {
		pair, ok := parensMap[chr]
		if ok {
			top, err := stack.Pop()
			if err != nil || top != pair {
				return false
			}
		} else {
			stack.Push(chr)
		}
	}
	return stack.Length == 0
}
