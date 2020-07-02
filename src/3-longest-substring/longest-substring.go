package longestSubstring

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lengthOfLongestSubstring(s string) int {
	var from, to, longest int
	seen := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		ch := s[i]
		_, exists := seen[ch]
		if exists {
			longest = max(longest, to-from)
			fromCh := s[from]
			for {
				delete(seen, fromCh)
				from++
				if fromCh == ch {
					break
				}
				fromCh = s[from]
			}
			seen[ch] = true
		} else {
			seen[ch] = true
		}
		to = i + 1
	}
	return max(longest, to-from)
}
