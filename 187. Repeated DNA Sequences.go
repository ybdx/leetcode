package leetcode

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}
	table := make(map[string]int, 0)
	res := make([]string, 0)

	for i := 0; i <= len(s)-10; i++ {
		temp := string(s[i : i+10])
		if value, ok := table[temp]; ok {
			table[temp] = value + 1
		} else {
			table[temp] = 1
		}
	}
	for k, v := range table {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}
