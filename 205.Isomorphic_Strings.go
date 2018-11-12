package leetcode

//func isIsomorphic(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//	hashMap1 := make(map[uint8]int)
//	hashMap2 := make(map[uint8]int)
//	for i:=0; i< len(s);i++ {
//		value1, _:= hashMap1[s[i]]
//		value2, _ := hashMap2[t[i]]
//		if value1 != value2 {
//			return false
//		} else {
//			hashMap1[s[i]] = i + 1
//			hashMap2[t[i]] = i + 1
//		}
//	}
//	return true
//}

// 4ms
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashMap1 := make(map[uint8]uint8)
	hashMap2 := make(map[uint8]uint8)
	for i:=0; i< len(s);i++ {
		if v, ok := hashMap1[s[i]]; ok {
			if v != t[i] {
				return false
			}
		} else if v, ok := hashMap2[t[i]]; ok {
			if v != s[i] {
				return false
			}
		} else {
			hashMap1[s[i]] = t[i]
			hashMap2[t[i]] = s[i]
		}
	}
	return true
}
