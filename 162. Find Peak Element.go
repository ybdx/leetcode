package leetcode

func findPeakElement(nums []int) int {
	if 0 == len(nums) {
		return -1
	}
	if 1 == len(nums) {
		return 0
	}
	i := 0
	for ; i < len(nums)-1; i++ {
		if i == 0 {
			if nums[i] > nums[i+1] {
				break
			}
		} else {
			if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
				break
			}
		}

	}
	return i
}


