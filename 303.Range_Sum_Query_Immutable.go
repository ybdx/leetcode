package leetcode

type NumArray struct {
	Nums []int
	Dp   []int
}

func Constructor(nums []int) NumArray {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if 0 == i {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}
	return NumArray{Nums: nums, Dp: dp}
}

func (this *NumArray) SumRange(i int, j int) int {
	if j >= len(this.Nums) {
		j = len(this.Nums) - 1
	}
	return this.Dp[j] - this.Dp[i] + this.Nums[i]
}
