package zigzag

func convert(s string, numRows int) string {
	var row, col int
	var result string
	sLen := len(s)
	diagonalLen := numRows - 2
	if diagonalLen < 0 {
		diagonalLen = 0
	}
	blockLen := numRows + diagonalLen
	blocksCount := sLen/blockLen + 1
	numColumns := blocksCount * (diagonalLen + 1)

	arr := make([][]rune, numRows)
	for i := range arr {
		arr[i] = make([]rune, numColumns)
	}

	for _, chr := range s {
		arr[row][col] = chr
		if numRows > 1 && col%(numRows-1) == 0 && row < (numRows-1) {
			// move down
			row++
		} else {
			// move by diagonal
			col++
			if row > 0 {
				row--
			}

		}
	}

	// collect chars
	for _, row := range arr {
		for _, chr := range row {
			if chr > 0 {
				result += string(chr)
			}
		}
	}

	return result
}
