package stringarrays

func unpack(word []string) []rune {
	var result []rune
	for _, wordPart := range word {
		for _, chr := range wordPart {
			result = append(result, chr)
		}
	}
	return result
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	word1Arr := unpack(word1)
	word2Arr := unpack(word2)

	if len(word1Arr) != len(word2Arr) {
		return false
	}

	for idx, chr := range word1Arr {
		if chr != word2Arr[idx] {
			return false
		}
	}

	return true
}
