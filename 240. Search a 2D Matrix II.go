package leetcode

func searchMatrix2(matrix [][]int, target int) bool {
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return false
	}
	row := 0
	col := len(matrix[0]) -1
	for row <= len(matrix) -1 && col >= 0 {
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			row ++
		} else {
			col --
		}
	}
	return false
}
