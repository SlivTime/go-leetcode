package twoSum

func twoSum(nums []int, target int) []int {
	var seen map[int]int
	seen = make(map[int]int)
	for i, val := range nums {
		complement := target - val
		j, ok := seen[complement]
		if ok {
			return []int{j, i}
		}
		seen[val] = i
	}
	return []int{}
}
