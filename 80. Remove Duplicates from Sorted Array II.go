package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dupCount := 1
	result := 1
	for i:=1; i<len(nums); i++ {
		if nums[i] == nums[i-1] {
			if dupCount < 2 {
				nums[result] = nums[i]
				result ++
			}
			dupCount++
		} else {
			dupCount = 1
			nums[result] = nums[i]
			result++
		}
	}
	return result
}