package longest_palindromic

// Ugly, but pretty effective (faster than 96.36% of Go submissions)
func longestPalindrome(s string) string {
	var cur, curLen, i, j int
	lenS := len(s)
	start := 0
	maxLen := 1

	for cur < len(s) {
		i = cur - 1
		j = cur + 1
		// odd
		for i >= 0 && j < lenS {
			if s[i] == s[j] {
				curLen = j - i + 1
				if curLen > maxLen {
					start = i
					maxLen = curLen
				}
				i = i - 1
				j = j + 1
			} else {
				break
			}
		}
		// even
		i = cur - 1
		j = cur
		for i >= 0 && j < lenS {
			if s[i] == s[j] {
				curLen = j - i + 1
				if curLen > maxLen {
					start = i
					maxLen = curLen
				}
				i = i - 1
				j = j + 1
			} else {
				break
			}
		}
		cur = cur + 1
	}

	return s[start : start+maxLen]
}
