package reverseinteger

import "math"

func reverse(x int) int {
	result := 0

	for x != 0 {
		result = result*10 + x%10
		if int(math.Abs(float64(result)))>>31 > 0 {
			return 0
		}
		x = x / 10
	}
	return result
}
