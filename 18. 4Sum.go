package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	results := make([][]int, 0)
	result := make([]int, 0)
	findNSum(nums, target, 4, result, &results)
	return results
}

// n sum
func findNSum(nums []int, target, n int, result []int, results *[][]int) {
	if n < 2 || len(nums) < n {
		return
	}
	if n == 2 {
		low := 0
		high := len(nums) - 1
		for low < high {
			if nums[low] + nums[high] > target {
				high --
			} else if nums[low] + nums[high] < target {
				low ++
			} else {
				temp := make([]int, len(result))
				copy(temp, result)
				temp = append(temp, nums[low], nums[high])
				*results = append(*results, temp)
				low ++
				high --
				for low < high && nums[low] == nums[low-1] {
					low ++
				}
				for low < high && nums[high] == nums[high + 1] {
					high --
				}
			}
		}
	} else {
		for i:=0; i<len(nums)-n+1; i++ {
			if i == 0 || (i > 0 && nums[i-1] != nums[i]) {
				temp := make([]int, len(result))
				copy(temp, result)
				temp = append(temp, nums[i])
				findNSum(nums[i+1:], target-nums[i], n-1, temp, results)
			}
		}
	}
}
