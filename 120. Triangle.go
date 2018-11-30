package leetcode

func minimumTotal(triangle [][]int) int {
	res := make([]int, len(triangle[len(triangle) - 1]) + 1)
	for i:= len(triangle) - 1; i>=0; i-- {
		for j:=0; j<len(triangle[i]); j++ {
			res[j] = min(res[j], res[j+1]) + triangle[i][j]
		}
	}
	return res[0]
}

