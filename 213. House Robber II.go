package leetcode

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res1 := []int {0, 0}
	for i:=0; i< len(nums)-1; i++ {
		temp:=res1[0]
		res1[0] = res1[1]
		res1[1] = max(temp + nums[i], res1[1])
	}
	res2 := []int{0, 0}
	for i:= 1; i < len(nums); i++ {
		temp:=res2[0]
		res2[0] = res2[1]
		res2[1] = max(temp + nums[i], res2[1])
	}
	return max(res1[1], res2[1])
}
