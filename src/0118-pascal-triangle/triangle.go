package triangle

func generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 0; j <= i; j++ {
			switch j {
			case 0, i:
				row[j] = 1
			default:
				row[j] = result[i-1][j-1] + result[i-1][j]
			}
		}
		result[i] = row
	}
	return result
}
