package longestsubstring

func lengthOfLongestSubstring(s string) int {
	var queue []rune
	var queueLen, maxLen, foundIndex int

	for _, chr := range s {
		foundIndex = -1
		// Find in queue
		for idx, qChr := range queue {
			if chr == qChr {
				foundIndex = idx
				break
			}
		}

		queue = append(queue, chr)
		if foundIndex >= 0 {
			queue = queue[foundIndex+1:]
		} else {
			// If not found, append and update len
			queueLen = len(queue)
			if queueLen > maxLen {
				maxLen = queueLen
			}
		}
	}

	return maxLen
}
