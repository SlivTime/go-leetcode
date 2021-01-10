package sortedarray

import "math"

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func findPosition(newValue int, nums []int) int {
	numsLen := len(nums)

	if numsLen == 0 {
		return 0
	}

	lo := 0
	hi := numsLen - 1

	for lo <= hi {
		mid := (lo + hi) / 2
		if newValue < nums[mid] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return lo
}

func createSortedArray(instructions []int) int {
	var position, cost int
	result := 0
	sortedNums := []int{}

	for _, num := range instructions {
		num = num * 2
		currentLen := len(sortedNums)
		loPosition := findPosition(num-1, sortedNums)
		hiPosition := findPosition(num+1, sortedNums)
		if loPosition < currentLen-hiPosition {
			position = loPosition
			cost = loPosition
		} else {
			position = hiPosition
			cost = currentLen - hiPosition
		}

		if position == currentLen {
			sortedNums = append(sortedNums, num)
		} else {
			sortedNums = append(sortedNums[:position+1], sortedNums[position:]...)
			sortedNums[position] = num
		}
		// fmt.Printf("Insert %5d at pos %d with cost %3d now nums = %v\n", num, position, cost, sortedNums)
		result += cost
	}
	return result % int(math.Pow(10, 9)+7)
}
