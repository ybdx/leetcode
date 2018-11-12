package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var count [26]int
	for i := 0; i < len(s); i++ {
		count[int(s[i]-'a')]++
		count[int(t[i]-'a')]--
	}
	for _, v := range count {
		if 0 != v {
			return false
		}
	}
	return true
}
