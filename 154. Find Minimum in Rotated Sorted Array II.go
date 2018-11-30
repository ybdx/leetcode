package leetcode

import "fmt"

func findMin2(nums []int) int {
	low := 0
	high := len(nums) - 1
	if len(nums) == 1 || nums[low] < nums[high] {
		return nums[low]
	}
	mid := low
	for nums[low] >= nums[high] {
		if high - low == 1 {
			mid = high
			break
		}
		mid = low + (high - low)/2
		if nums[low] == nums[high] && nums[mid] == nums[low] {
			mid = low

			for i := low + 1; i <=high; i++ {
				fmt.Println(mid, "===")
				if nums[mid] > nums[i] {
					fmt.Println("++++")
					mid = i
					fmt.Println(mid, "....")
				}
			}
			break
		} else if nums[mid] >= nums[low] {
			low = mid
		} else if nums[mid] <= nums[high] {
			high = mid
		}
	}
	fmt.Println(mid)
	return nums[mid]
}
