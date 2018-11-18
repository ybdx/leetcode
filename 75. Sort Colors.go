package leetcode

func sortColors(nums []int) {
	red, blue := 0, len(nums)-1
	for i := 0; i <= blue; i++ {
		for nums[i] == 2 && i < blue {
			nums[i], nums[blue] = nums[blue], nums[i]
			blue--
		}
		for nums[i] == 0 && red < i {
			nums[i], nums[red] = nums[red], nums[i]
			red++
		}
	}
}

