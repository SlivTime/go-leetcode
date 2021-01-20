package atoi

func myAtoi(s string) int {
	maxInt := 1<<31 - 1
	minInt := -1 << 31
	signSet := false

	result := 0
	sign := 1

	for _, chr := range s {
		if chr == ' ' && result == 0 && !signSet {
			continue
		} else if chr == '-' && !signSet {
			sign = -1
			signSet = true
		} else if chr == '+' && !signSet {
			signSet = true
		} else if chr >= '0' && chr <= '9' {
			result = result*10 + sign*int(chr-'0')
			signSet = true
		} else {
			break
		}

		if result > maxInt {
			result = maxInt
			break
		} else if result < minInt {
			result = minInt
			break
		}
	}
	return result
}
