package leetcode

// f(m, n) = f(m-1, n) + f(m, n-1)
// 采用dp做
func uniquePaths(m int, n int) int {
	path := make([][]int, n)
	for i:=0; i<n; i++ {
		temp := make([]int, m)
		temp[0] = 1
		path[i] = temp
	}
	for j:=0; j<m; j++ {
		path[0][j] = 1
	}
	for i :=1; i<n; i++ {
		for j:=1; j <m; j++ {
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}
	return path[n-1][m-1]
}
