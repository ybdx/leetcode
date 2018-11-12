package leetcode

func containsDuplicate(nums []int) bool {
	hashMap := make(map[int]int)
	for _, value := range nums {
		if _, ok := hashMap[value]; ok {
			return true
		}
		hashMap[value] = 1
	}
	return false
}
