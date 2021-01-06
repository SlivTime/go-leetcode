package find_number

func findKthPositive(arr []int, k int) int {
	missingNumber := 0

	for arrIndex, missingIndex := 0, 1; missingIndex <= k; {
		missingNumber++
		if arrIndex < len(arr) && missingNumber == arr[arrIndex] {
			arrIndex++
		} else {
			missingIndex++
		}
	}

	return missingNumber
}
