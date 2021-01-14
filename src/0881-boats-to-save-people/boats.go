package boats

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	if len(people) <= 1 {
		return len(people)
	}

	result := 0
	sort.Ints(people)
	lo := 0
	hi := len(people) - 1
	for lo <= hi {
		if lo == hi || people[lo]+people[hi] <= limit {
			result++
			lo++
			hi--
		} else {
			result++
			hi--
		}
	}
	return result
}
