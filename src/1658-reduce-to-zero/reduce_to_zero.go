package reducezero

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func minOperations(nums []int, x int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	target := totalSum - x
	sums := make(map[int]int)

	sums[0] = -1

	maxLenSubArray := -1
	currLenSubArray := 0
	prefixSum := 0

	for idx, num := range nums {
		prefixSum += num
		sums[prefixSum] = idx

		val, ok := sums[prefixSum-target]
		if ok {
			currLenSubArray = idx - val
			maxLenSubArray = max(maxLenSubArray, currLenSubArray)
		}
	}

	if maxLenSubArray == -1 {
		return maxLenSubArray
	} else {
		return len(nums) - maxLenSubArray
	}
}
