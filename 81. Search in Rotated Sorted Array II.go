package leetcode

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target || nums[low] == target {
			return true
		} else if nums[low] == nums[high] {
			low ++
			high --
		} else if nums[mid] > target {
			if nums[mid] > nums[high] && nums[low] > target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[mid] < nums[low] && nums[high] < target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return false
}
