package leetcode

func longestPalindrome(s string) int {
	result := 0
	temp := make([]int, 128)
	for i := 0; i < len(s); i++ {
		temp[int(s[i])] ++
	}
	for i := 0; i < 128; i++ {
		result += temp[i] / 2 * 2
		if result%2 == 0 && temp[i]%2 == 1 {
			result++
		}
	}
	return result
}
