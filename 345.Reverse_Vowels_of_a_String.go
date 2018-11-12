package leetcode

func reverseVowels(s string) string {
	// vowels is consist of a,e,i,o,u
	start := 0
	end := len(s) - 1
	source := []byte(s)
	for start < end {
		for ; start < end; start++ {
			if isVowels(source[start]) {
				break
			}
		}
		for ; end > start; end-- {
			if isVowels(source[end]) {
				break
			}
		}
		source[start], source[end] = source[end], source[start]
		start++
		end--
	}
	return string(source)
}

func isVowels(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'A' ||
		b == 'E' || b == 'I' || b == 'O' || b == 'U'
}
