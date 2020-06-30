package twoSum

func twoSum(nums []int, target int) []int {
	for i, val := range nums {
		for j := i + 1; j < len(nums); j++ {
			if val+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}
