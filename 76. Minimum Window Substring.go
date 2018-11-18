package leetcode

// 28ms
//func minWindow(s string, t string) string {
//	if len(s) < len(t) || len(s) == 0 || len(t) == 0 {
//		return ""
//	}
//	tMap := make(map[byte]int)
//	for i := 0; i < len(t); i++ {
//		if v, ok := tMap[t[i]]; ok {
//			tMap[t[i]] = v + 1
//		} else {
//			tMap[t[i]] = 1
//		}
//	}
//	required := len(tMap)
//	//ans list of the form (window length, left, right)
//	ans := []int{-1, 0, 0}
//	windowMap := make(map[byte]int)
//	l, r, formed := 0, 0, 0
//	for r < len(s) {
//		sv, ok := windowMap[s[r]]
//		if ok {
//			windowMap[s[r]] = sv + 1
//		} else {
//			windowMap[s[r]] = 1
//		}
//		tv, ok1 := tMap[s[r]]
//		if ok1 && ok && tv == sv+1 || ok1 && !ok && tv == 1 {
//			formed ++
//		}
//		for l <= r && formed == required {
//			if ans[0] == -1 || r-l+1 < ans[0] {
//				ans[0] = r - l + 1
//				ans[1] = l
//				ans[2] = r
//			}
//			sv, ok = windowMap[s[l]]
//			windowMap[s[l]] --
//			tv, ok1 = tMap[s[l]]
//			if ok1 && sv-1 < tv {
//				formed --
//			}
//			l++
//		}
//		r++
//	}
//	if ans[0] == -1 {
//		return ""
//	}
//	return s[ans[1] : ans[2]+1]
//}

// 4ms
func minWindow(s string, t string) string {
	if len(s) < len(t) || len(s) == 0 || len(t) == 0 {
		return ""
	}
	window := make([]int, 256)
	need := make([]int, 256)
	start, end := 0, 0
	formed, required := 0, len(t)
	minValue := len(s) + 1
	res := ""
	for i := 0; i < len(t); i++ {
		need[t[i]] ++
	}
	for end < len(s) {
		if need[s[end]] > window[s[end]] {
			formed ++
		}
		window[s[end]] ++
		for start <= end && window[s[start]] > need[s[start]] {
			window[s[start]] --
			start++
		}
		if formed == required && (end-start+1) < minValue {
			minValue = end - start + 1
			res = s[start : end+1]
		}
		end ++
	}
	return res
}
