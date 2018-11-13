package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	res := make([]int, 128)
	for i := 0; i < len(magazine); i++ {
		res[int(magazine[i])] ++
	}
	for i := 0; i < len(ransomNote); i++ {
		res[int(ransomNote[i])] --
		if res[int(ransomNote[i])] < 0 {
			return false
		}
	}
	return true
}
