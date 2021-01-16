package getmax

func isEven(n int) bool {
	return n%2 == 0
}

func getMaximumGenerated(n int) int {
	var maxValue int
	cache := make(map[int]int)
	cache[0] = 0
	cache[1] = 1
	if n > 0 {
		maxValue = cache[1]
	}
	for i := 2; i <= n; i++ {
		if isEven(i) {
			cache[i] = cache[i/2]
		} else {
			cache[i] = cache[i/2] + cache[i/2+1]
		}
		if cache[i] > maxValue {
			maxValue = cache[i]
		}
	}
	return maxValue
}
