package sumpairs

func inc(pairs map[int]int, key int) {
	val, ok := pairs[key]
	if ok {
		pairs[key] = val + 1
	} else {
		pairs[key] = 1
	}
}

func dec(pairs map[int]int, key int) {
	val, ok := pairs[key]
	if !ok {
		panic("No value for this key")
	}
	if val > 1 {
		pairs[key] = val - 1
	} else {
		delete(pairs, key)
	}
}

func maxOperations(nums []int, k int) int {
	count := len(nums)
	finalCount := 0
	pairs := make(map[int]int)

	for _, val := range nums {
		_, ok := pairs[k-val]
		if ok {
			dec(pairs, k-val)
		} else {
			inc(pairs, val)
		}
	}

	for _, v := range pairs {
		finalCount += v
	}

	return (count - finalCount) / 2
}
