package wordladder

func wordDifference(beginWord string, endWord string) int {
	diff := 0
	for idx := range beginWord {
		if beginWord[idx] != endWord[idx] {
			diff++
		}
	}
	return diff
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var distance int
	var pointsToVisit []int

	paths := make(map[int][]int)
	fullList := append([]string{beginWord}, wordList...)
	wordsLen := len(fullList)

	distances := make([]int, wordsLen)
	for idx := range distances {
		if idx > 0 {
			distances[idx] = wordsLen
		}
	}
	seen := make(map[int]bool, wordsLen)

	// init paths and distances
	for currIdx, currentWord := range fullList {
		for otherIdx, otherWord := range fullList {
			distance = wordDifference(currentWord, otherWord)
			if distance == 1 {
				paths[currIdx] = append(paths[currIdx], otherIdx)
			}
		}
	}

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
				neighbors := paths[current]
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
