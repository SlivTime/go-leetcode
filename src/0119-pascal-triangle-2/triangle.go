package triangle2

type Key struct {
	col int
	row int
}

var cache = make(map[Key]int)

func getItem(rowIndex int, colIndex int) int {
	key := Key{colIndex, rowIndex}
	cached, ok := cache[key]
	if ok {
		return cached
	} else {
		if colIndex == 0 || rowIndex == colIndex {
			cached = 1
		} else {
			cached = getItem(rowIndex-1, colIndex-1) + getItem(rowIndex-1, colIndex)
		}
		cache[key] = cached
	}

	return cached
}

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	for colIndex := 0; colIndex <= rowIndex; colIndex++ {
		result[colIndex] = getItem(rowIndex, colIndex)
	}
	return result
}
