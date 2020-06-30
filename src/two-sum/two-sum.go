package twoSum

func twoSum(nums []int, target int) []int {
	for i, val := range nums {
		for j := i + 1; j < len(nums); j++ {
			sum := val + nums[j]
			if sum == target {
				return []int{i, j}
			}
			if sum > target {
				continue
			}
		}
	}
	return []int{0, 0}
}
