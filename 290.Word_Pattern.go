package leetcode

import "strings"

func wordPattern(pattern string, str string) bool {
	arr := strings.Split(str, " ")
	if len(pattern) != len(arr) {
		return false
	}
	map1 := make(map[byte]string)
	map2 := make(map[string]byte)
	for i:=0; i < len(pattern); i++ {
		if v, ok := map1[pattern[i]]; ok {
			if v != arr[i] {
				return false
			}
		} else if v, ok:= map2[arr[i]]; ok {
			if v != pattern[i] {
				return false
			}
		} else {
			map1[pattern[i]] = arr[i]
			map2[arr[i]] = pattern[i]
		}
	}
	return true
}
