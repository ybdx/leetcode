package leetcode

//func isInterleave(s1 string, s2 string, s3 string) bool {
//	fmt.Println(s1, s2, s3)
//	if s1 != "" && s2 != "" && s1[0] == s3[0] && s2[0] != s3[0] {
//		return isInterleave(s1[1:], s2, s3[1:])
//	} else if s1 != "" && s2 != "" && s1[0] != s3[0] && s2[0] == s3[0] {
//		return isInterleave(s1, s2[1:], s3[1:])
//	} else if s1 != "" && s2 != "" && s1[0] == s3[0] && s2[0] == s3[0] {
//		return isInterleave(s1, s2[1:], s3[1:]) || isInterleave(s1[1:], s2, s3[1:])
//	} else if s1 == "" && s2 == s3  || s2 == "" && s1 == s3 || s1 == "" && s2 == "" && s3 == "" {
//		return true
//	}
//	return false
//}


func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) {
		return false
	}
	dp := make([][]bool, 0)
	for i:=0; i <= len(s1); i++ {
		temp := make([]bool, len(s2) + 1)
		dp = append(dp, temp)
	}

	for i:=0; i<=len(s1); i++ {
		for j:=0; j <=len(s2); j++ {
			if 0 == i && 0 == j {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			} else {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1] || dp[i-1][j] && s1[i-1] == s3[i+j-1]
			}
		}
	}
	return dp[len(s1)][len(s2)]
}