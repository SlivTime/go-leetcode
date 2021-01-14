package wordladder

func isNeughbor(beginWord string, endWord string) bool {
	diff := 0
	for idx := range beginWord {
		if beginWord[idx] != endWord[idx] {
			diff++
		}
	}
	return diff == 1
}

func getNeighbors(word string, wordList []string) []int {
	var result []int
	for idx, otherWord := range wordList {
		if isNeughbor(word, otherWord) {
			result = append(result, idx)
		}
	}
	return result
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var distance int
	var pointsToVisit []int

	fullList := append([]string{beginWord}, wordList...)
	wordsLen := len(fullList)

	distances := make([]int, wordsLen)
	for idx := range distances {
		if idx > 0 {
			distances[idx] = wordsLen
		}
	}
	seen := make(map[int]bool, wordsLen)

	for {
		pointsToVisit = pointsToVisit[:0]
		for idx, nextDistance := range distances {
			if nextDistance == distance && !seen[idx] {
				pointsToVisit = append(pointsToVisit, idx)
			}
		}

		if len(pointsToVisit) == 0 {
			break
		} else {
			for _, current := range pointsToVisit {
				neighbors := getNeighbors(fullList[current], fullList)
				seen[current] = true
				if fullList[current] == endWord {
					return distance + 1
				}
				for _, nextIdx := range neighbors {
					if !seen[nextIdx] {
						distances[nextIdx] = distance + 1
					}
				}
			}
		}
		distance++
	}

	return 0
}
