package leetcode

func thirdMax(nums []int) int {
	MIN := ^(1 << 32) + 1
	firstMax := MIN
	secondMax := MIN
	thirdMax := MIN
	for i := 0; i < len(nums); i++ {
		if nums[i] == firstMax || nums[i] == secondMax || nums[i] == thirdMax {
			continue
		}
		if nums[i] > firstMax {
			thirdMax = secondMax
			secondMax = firstMax
			firstMax = nums[i]
		} else if nums[i] > secondMax {
			thirdMax = secondMax
			secondMax = nums[i]
		} else if nums[i] > thirdMax {
			thirdMax = nums[i]
		}
	}
	if thirdMax == MIN {
		return firstMax
	}
	return thirdMax
}
