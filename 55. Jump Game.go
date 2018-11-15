package leetcode

// from start to end
func canJump(nums []int) bool {
	lastPos := 0
	for i:=0; i<len(nums) - 1; i++ {
		if nums[i] + i > lastPos && i <= lastPos {
			lastPos = nums[i] + i
		}
		if lastPos >= len(nums) - 1 {
			return true
		}
	}
	return lastPos >= len(nums) - 1
}


// from end to start
func canJump1(nums []int) bool {
	lastPos := len(nums) - 1
	for i:= lastPos; i >=0; i-- {
		if i + nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}
