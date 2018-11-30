package leetcode

func findMin(nums []int) int {
	low := 0
	high := len(nums) - 1
	if nums[low] <= nums[high] {
		return nums[low]
	}
	mid := high
	for nums[low] >= nums[high] {
		if high - low == 1 {
			mid = high
			break
		}
		mid = low + (high -low)/2
		if nums[mid] >= nums[high] {
			low = mid
		} else if  nums[mid] <=  nums[low]  {
			high = mid
		}
	}
	return nums[mid]
}
