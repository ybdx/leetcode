package leetcode

func firstUniqChar(s string) int {
	arr := make([]int, 128)
	for i := 0; i < len(s); i++ {
		arr[int(s[i])] ++
	}
	for i := 0; i < len(s); i++ {
		if arr[int(s[i])] == 1 {
			return i
		}
	}
	return -1
}
