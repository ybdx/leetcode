package leetcode

func longestConsecutive(nums []int) int {
	hashMap := make(map[int]int)
	for i:=0; i<len(nums); i++ {
		hashMap[nums[i]] = 1
	}
	result := 0
	for i:=0; i<len(nums); i++ {
		if _, ok:= hashMap[nums[i] - 1]; !ok {
			curLength := 1
			curNumber := nums[i]
			_, ok := hashMap[curNumber + 1]
			for ok {
				curNumber++
				curLength++
				_, ok = hashMap[curNumber + 1]
			}
			result = max(result, curLength)
		}
	}
	return result
}

