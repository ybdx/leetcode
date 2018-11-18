package leetcode

//Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the
// sum of all numbers along its path.
//
//Note: You can only move either down or right at any point in time.
//
//Example:
//
//Input:
//[
//[1,3,1],
//[1,5,1],
//[4,2,1]
//]
//Output: 7
//Explanation: Because the path 1→3→1→1→1 minimizes the sum.

// f(m, n) = min(f(m-1, n), f(m, n-1)) + curNum
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	for i:=1; i<m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i:=1; i<n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i:=1; i<m; i++ {
		for j:=1; j<n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}
