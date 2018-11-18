package leetcode

// 可以分为三种情况：
// 1.一步操作之后，再将A[2, ..., lenA], B[1, ..., lenB]变成相同字符串
// 2.一步操作之后，再将A[1, ..., lenA], B[2, ..., lenB]变成相同字符串
// 3.一步操作之后，再将A[2, ..., lenA], B[2, ..., lenB]变成相同字符串
// 公式：f(m, n) = min{f(m-1, n), f(m, n-1), f(m-1, n-1)} + 1

// 递归实现，重复计算的太多 复杂度高
//func minDistance(word1 string, word2 string) int {
//	return calculateStringDistance(word1, word2, 0, len(word1), 0, len(word2))
//}
//
//func calculateStringDistance(word1 string, word2 string, w1Start, w1End, w2Start, w2End int) int {
//	if w1Start >= w1End {
//		if w2Start >= w2End {
//			return 0
//		} else {
//			return w2End - w2Start
//		}
//	}
//	if w2Start >= w2End {
//		if w1Start >= w1End {
//			return 0
//		} else {
//			return w1End - w1Start
//		}
//	}
//	if word1[w1Start] == word2[w2Start] {
//		return calculateStringDistance(word1, word2, w1Start+1, w1End, w2Start+1, w2End)
//	} else {
//		t1 := calculateStringDistance(word1, word2, w1Start+1, w1End, w2Start+1, w2End)
//		t2 := calculateStringDistance(word1, word2, w1Start, w1End, w2Start+1, w2End)
//		t3 := calculateStringDistance(word1, word2, w1Start+1, w1End, w2Start, w2End)
//		return min(min(t1, t2), t3) + 1
//	}
//}

// 非递归，不重复计算
func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	m := len(word1)
	n := len(word2)
	cost := make([][]int, 0)
	for i := 0; i <= m; i++ {
		temp := make([]int, n+1)
		cost = append(cost, temp)
	}
	for i := 0; i <= m; i++ {
		cost[i][0] = i
	}
	for j := 0; j <= n; j++ {
		cost[0][j] = j
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				cost[i+1][j+1] = cost[i][j]
			} else {
				t1 := cost[i][j]
				t2 := cost[i+1][j]
				t3 := cost[i][j+1]
				cost[i+1][j+1] = min(min(t1, t2), t3) + 1
			}
		}
	}
	return cost[m][n]
}
