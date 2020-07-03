package median

const maxInt = int(^uint(0) >> 1)

func safeGetByIndex(arr []int, idx int) int {
	if idx < len(arr) {
		return arr[idx]
	}
	return maxInt
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	resultArray := make([]int, len(nums1)+len(nums2))
	idx1 := 0
	idx2 := 0
	item1 := maxInt
	item2 := maxInt
	midIndex := cap(resultArray) / 2.0
	var median float64
	resultIsOdd := cap(resultArray)%2 == 1

	for i := 0; i < cap(resultArray); i++ {
		item1 = safeGetByIndex(nums1, idx1)
		item2 = safeGetByIndex(nums2, idx2)

		if item1 < item2 {
			resultArray[i] = item1
			idx1++
		} else {
			resultArray[i] = item2
			idx2++
		}

	}

	if resultIsOdd {
		median = float64(resultArray[len(resultArray)/2])
	} else {
		median = float64(float64(resultArray[midIndex]+resultArray[midIndex-1]) / 2)
	}

	return median
}
